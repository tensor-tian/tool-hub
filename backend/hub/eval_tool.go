package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"tool-hub/backend/hub/fifo"

	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// EvalToolRequestEvent represents the event sent to frontend for tool evaluation
type EvalToolRequestEvent struct {
	RequestID  string `json:"requestId"`
	Code       string `json:"code"`
	Parameters string `json:"parameters"`
}

// EvalToolResponseEvent represents the event received from frontend after tool evaluation
type EvalToolResponseEvent struct {
	RequestID string          `json:"requestId"`
	Success   bool            `json:"success"`
	Tool      json.RawMessage `json:"tool"`
	Error     string          `json:"error"`
}

var (
	evalToolLimiter       = fifo.DefaultGroupLimiter
	pendingEvalRequests   = make(map[string]chan EvalToolResponseEvent)
	pendingEvalRequestsMu sync.RWMutex
)

// InitToolEvalListener sets up the global event listener for tool evaluation responses
func InitToolEvalListener(ctx context.Context) {
	runtime.EventsOn(ctx, "eval-tool-response", func(data ...interface{}) {
		if len(data) > 0 {
			if responseMap, ok := data[0].(map[string]interface{}); ok {
				var response EvalToolResponseEvent
				// Convert map to struct
				jsonBytes, _ := json.Marshal(responseMap)
				json.Unmarshal(jsonBytes, &response)

				// Find the corresponding pending request
				pendingEvalRequestsMu.RLock()
				responseChan, exists := pendingEvalRequests[response.RequestID]
				pendingEvalRequestsMu.RUnlock()

				if exists {
					select {
					case responseChan <- response:
					default:
					}
				}
			}
		}
	})
}

// EvalTool evaluates a tool plugin with the given code and parameters using the frontend WebWorker
func EvalTool(ctx context.Context, code string, parameters string) (json.RawMessage, error) {
	// Acquire semaphore to prevent event confusion (only one eval at a time)
	if err := evalToolLimiter.Acquire(ctx, "eval-tool", 1); err != nil {
		return nil, fmt.Errorf("failed to acquire eval lock: %w", err)
	}

	// Generate unique request ID
	requestID := uuid.New().String()

	// Create channel to receive response
	responseChan := make(chan EvalToolResponseEvent, 1)

	// Register this request
	pendingEvalRequestsMu.Lock()
	pendingEvalRequests[requestID] = responseChan
	pendingEvalRequestsMu.Unlock()

	// Clean up after we're done
	defer func() {
		pendingEvalRequestsMu.Lock()
		delete(pendingEvalRequests, requestID)
		pendingEvalRequestsMu.Unlock()
	}()

	// Emit request to frontend
	request := EvalToolRequestEvent{
		RequestID:  requestID,
		Code:       code,
		Parameters: parameters,
	}
	runtime.EventsEmit(ctx, "eval-tool-request", request)

	// Wait for response with timeout
	var response EvalToolResponseEvent
	select {
	case response = <-responseChan:
		// Release the limiter immediately after receiving response
		evalToolLimiter.Release("eval-tool")

	case <-time.After(30 * time.Second):
		evalToolLimiter.Release("eval-tool")
		return nil, fmt.Errorf("tool evaluation timeout")
	}

	if !response.Success {
		return nil, fmt.Errorf("tool evaluation failed: %s", response.Error)
	}

	return response.Tool, nil
}
