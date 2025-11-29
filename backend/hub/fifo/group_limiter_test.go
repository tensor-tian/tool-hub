package fifo

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGroupLimiter_GetSemaphoreFor(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(5)

	sem1 := m.getSemaphoreFor(groupName, maxConcurrent)
	assert.NotNil(t, sem1)

	// Should return same semaphore for same group
	sem2 := m.getSemaphoreFor(groupName, maxConcurrent)
	assert.Equal(t, sem1, sem2)
}

func TestGroupLimiter_GetSemaphoreFor_ZeroMaxConcurrent(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(0)

	sem := m.getSemaphoreFor(groupName, maxConcurrent)
	assert.NotNil(t, sem)

	// Try to acquire once (should succeed), and a second time (should block)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := sem.Acquire(ctx)
	assert.NoError(t, err, "should be able to acquire the only permit")

	done := make(chan error)
	go func() {
		ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel2()
		err := sem.Acquire(ctx2)
		done <- err
	}()

	select {
	case err := <-done:
		// Should timeout because semaphore is full
		assert.Error(t, err, "should timeout when trying to acquire a second permit when max concurrent is 1")
	case <-time.After(200 * time.Millisecond):
		t.Error("test timed out waiting for second acquire to fail")
	}
	sem.Release()
}

func TestGroupLimiter_AcquireRelease(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(2)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Acquire first permit
	err := m.Acquire(ctx, groupName, maxConcurrent)
	assert.NoError(t, err)

	// Acquire second permit
	err = m.Acquire(ctx, groupName, maxConcurrent)
	assert.NoError(t, err)

	// Release permits
	m.Release(groupName)
	m.Release(groupName)
}

func TestGroupLimiter_AcquireBlocking(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Acquire the only permit
	err := m.Acquire(ctx, groupName, maxConcurrent)
	assert.NoError(t, err)

	acquired := false
	var wg sync.WaitGroup
	wg.Add(1)

	// Try to acquire in goroutine - should block
	go func() {
		defer wg.Done()
		ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel2()
		err := m.Acquire(ctx2, groupName, maxConcurrent)
		assert.NoError(t, err)
		acquired = true
	}()

	// Give goroutine time to block
	time.Sleep(100 * time.Millisecond)
	assert.False(t, acquired, "should be blocking")

	// Release to unblock
	m.Release(groupName)

	// Wait for goroutine to complete with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		assert.True(t, acquired, "should have acquired after release")
	case <-time.After(2 * time.Second):
		t.Fatal("test timed out waiting for acquire")
	}
}

func TestGroupLimiter_ConcurrentOperations(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(5)

	const numGoroutines = 20
	var wg sync.WaitGroup
	successCount := int32(0)
	var mu sync.Mutex

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := m.Acquire(ctx, groupName, maxConcurrent)
			if err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
				time.Sleep(10 * time.Millisecond)
				m.Release(groupName)
			}
		}()
	}

	// Wait for all goroutines to complete with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		assert.Equal(t, int32(numGoroutines), successCount)
	case <-time.After(15 * time.Second):
		t.Fatal("test timed out waiting for concurrent operations")
	}
}

func TestGroupLimiter_ReleaseNonExistentGroup(t *testing.T) {
	m := NewGroupLimiter()

	// Should not panic
	assert.NotPanics(t, func() {
		m.Release("999")
	})
}

func TestGroupLimiter_Reset(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(5)

	sem := m.getSemaphoreFor(groupName, maxConcurrent)
	assert.NotNil(t, sem)

	// Store the pointer address
	oldSemPtr := sem

	// Reset should remove the semaphore
	m.Reset(groupName)

	m.mu.RLock()
	_, exists := m.groups[groupName]
	m.mu.RUnlock()
	assert.False(t, exists, "semaphore should be removed from groups map")

	// Next getSemaphoreFor should create a new one
	newSem := m.getSemaphoreFor(groupName, maxConcurrent)
	assert.NotNil(t, newSem)

	// Verify it's a different instance by comparing pointers
	assert.NotSame(t, oldSemPtr, newSem, "should create a new semaphore instance")
}

func TestGroupLimiter_ConcurrentResetAndAcquire(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(3)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	// Concurrently acquire, reset, and release
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Try to acquire
			err := m.Acquire(ctx, groupName, maxConcurrent)
			if err == nil {
				// Simulate some work
				time.Sleep(5 * time.Millisecond)
				m.Release(groupName)
			}

			// Occasionally reset
			if i%3 == 0 {
				m.Reset(groupName)
			}
		}(i)
	}

	// Wait for all goroutines to complete with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Should not panic or deadlock
	case <-time.After(15 * time.Second):
		t.Fatal("test timed out waiting for concurrent reset and acquire")
	}
}

func TestGroupLimiter_ContextCancellation(t *testing.T) {
	m := NewGroupLimiter()

	groupName := "1"
	maxConcurrent := uint(1)

	// First acquire the only permit
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1()
	err := m.Acquire(ctx1, groupName, maxConcurrent)
	assert.NoError(t, err)

	// Try to acquire with a cancelled context
	ctx2, cancel2 := context.WithCancel(context.Background())
	cancel2() // Cancel immediately

	err = m.Acquire(ctx2, groupName, maxConcurrent)
	assert.Error(t, err, "should return error when context is cancelled")

	// Clean up
	m.Release(groupName)
}

func TestGroupLimiter_MultipleGroupsConcurrency(t *testing.T) {
	m := NewGroupLimiter()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	numGroups := 5
	numGoroutinesPerGroup := 10

	for groupID := uint(1); groupID <= uint(numGroups); groupID++ {
		for i := 0; i < numGoroutinesPerGroup; i++ {
			wg.Add(1)
			go func(gName string) {
				defer wg.Done()
				err := m.Acquire(ctx, gName, 3)
				if err == nil {
					time.Sleep(5 * time.Millisecond)
					m.Release(gName)
				}
			}(fmt.Sprintf("%d", groupID))
		}
	}

	// Wait for all goroutines to complete with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// All operations completed successfully
	case <-time.After(15 * time.Second):
		t.Fatal("test timed out waiting for multiple groups concurrency")
	}
}
