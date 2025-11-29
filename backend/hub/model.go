package hub

import (
	"encoding/json"
	"fmt"
)

type IUpdate interface {
	getUpdatedAt() int64
}

type ICreate interface {
	getBaseModel() BaseModel
}

// BaseModel provides common fields for all database models.
type BaseModel struct {
	ID        int   `json:"id" gorm:"primarykey"`
	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}

type ConcurrencyGroupName string

const (
	ConcurrencyGroupSystem ConcurrencyGroupName = "system"
	ConcurrencyGroupTool   ConcurrencyGroupName = "tool"
	ConcurrencyGroupNone   ConcurrencyGroupName = "none"
)

// ConcurrencyGroup represents a group of tools with concurrency control.
type ConcurrencyGroup struct {
	BaseModel
	Name          ConcurrencyGroupName `json:"name" gorm:"uniqueIndex"`
	Description   string               `json:"description"`
	MaxConcurrent uint                 `json:"maxConcurrent"` // max concurrent tasks in the group
}

func (g ConcurrencyGroup) GroupName(toolID int) (isValid bool, name string) {
	if g.Name == "" || g.Name == ConcurrencyGroupNone {
		return false, ""
	}
	if g.Name == ConcurrencyGroupTool {
		return true, fmt.Sprintf("%s_%d", g.Name, toolID)
	}
	return true, string(g.Name)
}

func (c ConcurrencyGroup) getUpdatedAt() int64 {
	return c.UpdatedAt
}

func (c ConcurrencyGroup) getBaseModel() BaseModel {
	return c.BaseModel
}

type Tool struct {
	BaseModel
	Name                 string `json:"name" gorm:"uniqueIndex"`
	Description          string `json:"description"`
	Parameters           string `json:"parameters"`           // JSON schema for parameters
	Type                 string `json:"type"`                 // "command_line", "http", "service"
	LogLifeSpan          string `json:"logLifeSpan"`          // log life span e.g. "24h", "7d"
	ConcurrencyGroupName string `json:"concurrencyGroupName"` // ConcurrencyGroup.Name
}

func (c Tool) getUpdatedAt() int64 {
	return c.UpdatedAt
}

func (c Tool) getBaseModel() BaseModel {
	return c.BaseModel
}

// StringMatcher represents a single matcher used to validate a string output;
type StringMatchType string

const (
	MatcherContains StringMatchType = "contains"
	MatcherEquals   StringMatchType = "equals"
	MatcherRegex    StringMatchType = "regex"
	MatcherPrefix   StringMatchType = "prefix"
	MatcherSuffix   StringMatchType = "suffix"
)

// ToolTestcase represents a test case for a tool invocation.
type ToolTestcase struct {
	BaseModel
	ToolID    int             `json:"toolID"`
	Args      string          `json:"args" `
	MatchType StringMatchType `json:"matchType"`
	Expected  string          `json:"expected"`
	Passed    bool            `json:"passed"`
}

func (ttc ToolTestcase) getUpdatedAt() int64 {
	return ttc.UpdatedAt
}

func (ttc ToolTestcase) getBaseModel() BaseModel {
	return ttc.BaseModel
}

type StringMatcher struct{}

// CommandLineTool represents a command line tool that can be invoked via HTTP POST request
type CommandLineTool struct {
	BaseModel
	Sh       string `json:"sh"`
	WD       string `json:"wd"`
	Cmd      string `json:"cmd"` // "sqlite3 ./hub.db \"SELECT sql FROM sqlite_master WHERE type='table' AND name='dependencies'\"| pg_format > hub.sql"
	Env      string `json:"env"` // environment variables in JSON format
	Timeout  string `json:"timeout"`
	IsStream bool   `json:"isStream"`
}

func (clt CommandLineTool) getUpdatedAt() int64 {
	return clt.UpdatedAt
}

func (clt CommandLineTool) getBaseModel() BaseModel {
	return clt.BaseModel
}

// HTTPTool represents a tool that makes HTTP calls to external endpoints.
type HTTPTool struct {
	BaseModel
	URL     string `json:"url"`     // HTTP URL
	Method  string `json:"method"`  // HTTP method e.g. "POST", "GET"
	Query   string `json:"query"`   // URL query parameters in JSON format
	Headers string `json:"headers"` // HTTP headers in JSON format
	Body    string `json:"body"`    // request body template
	Timeout string `json:"timeout"` // e.g. "30s"
}

type HTTPTestcase struct {
	BaseModel
	ToolID int    `json:"toolID"`
	Args   string `json:"args"` // input arguments in JSON format
}

// ServiceTool represents a service-based tool with its status and error information.
type ServiceTool struct {
	BaseModel
	StartCmd string `json:"startCmd"` // command to start the service
	Error    string `json:"error"`    // last error message
	Status   string `json:"status"`   // "active", "error", "ready"
}

// CallingLog represents a record of a tool invocation
// caller calls callee with input and output
type CallingLog struct {
	BaseModel
	CallerID   *int    `json:"callerID"`   // nullable, if null means external caller, caller is not in our ecosystem
	CallerType *string `json:"callerType"` // "http", "command", "service"
	CalleeID   int     `json:"calleeID"`
	CalleeType string  `json:"calleeType"`
	Input      string  `json:"input"`
	Output     string  `json:"output"`
	Error      string  `json:"error"`
	Duration   string  `json:"duration"`
	ExpiredAt  string  `json:"expiredAt"`
}

func fromMap[T any](m map[string]any) (T, error) {
	var result T
	bs, err := json.Marshal(m)
	if err != nil {
		return result, err
	}
	json.Unmarshal(bs, &result)
	return result, nil
}
