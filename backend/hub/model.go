package hub

import (
	"encoding/json"
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

// #region Tool

// Tool represents a tool stored in db.
// db schema
type Tool struct {
	BaseModel
	Name          string `json:"name" gorm:"uniqueIndex"`
	Description   string `json:"description"`
	Parameters    string `json:"parameters"` // json schema of parameters
	Category      string `json:"category"`
	Schema        string `json:"schema"`        // serialized zod schema of parameters for validation
	Definition    string `json:"definition"`    // typescript definition of parameters
	Code          string `json:"code"`          // plugin code
	DefaultParams string `json:"defaultParams"` // default parameters in json format
}

type CategoryOfTool string

const (
	CategoryCommandLine CategoryOfTool = "commandLine"
	CategoryHTTP        CategoryOfTool = "http"
)

// ToolCategory represents a category of tools.
// not db schema
type ToolCategory struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

// CommandLineToolExtra represents detailed info of a command line tool.
// matches with tool-hub-cli/utils
// not db schema
type CommandLineToolExtra struct {
	Sh    string `json:"sh"`
	WD    string `json:"wd"`
	Cmd   string `json:"cmd"` // "sqlite3 ./hub.db \"SELECT sql FROM sqlite_master WHERE type='table' AND name='dependencies'\"| pg_format > hub.sql"
	Env   string `json:"env"` // environment variables in JSON format
	Stdin string `json:"stdin"`
}

// CommandLineTool represents full info of a command line tool.
// matches with tool-hub-cli/utils
// not db schema
type CommandLineTool struct {
	Tool
	LogLifeSpan          string               `json:"logLifeSpan"`          // log life span e.g. "24h", "7d"
	ConcurrencyGroupName string               `json:"concurrencyGroupName"` // ConcurrencyGroup.Name
	Timeout              string               `json:"timeout"`
	IsStream             bool                 `json:"isStream"`
	Extra                CommandLineToolExtra `json:"extra"` // extra settings
}

// matches with tool-hub-cli/utils
// not db schema
type HTTPToolExtra struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Query   string            `json:"query"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// matches with tool-hub-cli/utils
// not db schema
type HTTPTool struct {
	Tool
	LogLifeSpan          string        `json:"logLifeSpan"`          // log life span e.g. "24h", "7d"
	ConcurrencyGroupName string        `json:"concurrencyGroupName"` // ConcurrencyGroup.Name
	Timeout              string        `json:"timeout"`
	IsStream             bool          `json:"isStream"`
	Extra                HTTPToolExtra `json:"extra"` // extra settings
}

// #endregion

// Setting represents a key-value setting of tool-hub app stored in db.
// db schema
type Setting struct {
	Key   string `json:"key" gorm:"primarykey"`
	Value string `json:"value"`
}

// ToolTestcase represents a testcase for a tool.
// db schema
type ToolTestcase struct {
	BaseModel
	ToolName string `json:"toolName"`
	Input    string `json:"input"`
	Output   string `json:"output"`
	OK       bool   `json:"ok"`
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

// ParseToolByCategory parses a JSON string into the appropriate tool type based on category
func ParseToolByCategory(toolJSON string, category string) (interface{}, error) {
	switch CategoryOfTool(category) {
	case CategoryCommandLine:
		var tool CommandLineTool
		if err := json.Unmarshal([]byte(toolJSON), &tool); err != nil {
			return nil, err
		}
		return tool, nil
	case CategoryHTTP:
		var tool HTTPTool
		if err := json.Unmarshal([]byte(toolJSON), &tool); err != nil {
			return nil, err
		}
		return tool, nil
	default:
		var tool Tool
		if err := json.Unmarshal([]byte(toolJSON), &tool); err != nil {
			return nil, err
		}
		return tool, nil
	}
}
