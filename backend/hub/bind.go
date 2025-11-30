package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

// Model is the main struct for tool hub operations.
type Model struct {
	ctx     context.Context
	HomeDir string `json:"homeDir"`
}

var model Model = func() Model {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}
	return Model{
		HomeDir: homeDir,
	}
}()

// ExposeModels returns the singleton instance of api.
func ExposeModels() []any {
	return []any{&model}
}

func setModelContext(ctx context.Context) {
	model.ctx = ctx
}

// #region Constant

// Dirs represents a map of directory paths.
type Dirs struct {
	Home string `json:"home"`
	Temp string `json:"temp"`
	App  string `json:"app"`
}

// GetDirs returns the Local directories path.
func (a *Model) GetDirs() Dirs {
	home, err := os.UserHomeDir()
	if err != nil {
		home = ""
	}
	temp := os.TempDir()
	return Dirs{
		Home: home,
		Temp: temp,
		App:  path.Join(home, ".config/tool-hub"),
	}
}

// #endregion

// #region Settings

type RespGetSettings struct {
	Error string            `json:"error"`
	KVMap map[string]string `json:"kvMap"`
}

func (m *Model) GetSettings(keys []string) (resp RespGetSettings) {
	resp.KVMap = make(map[string]string)
	list, err := gorm.G[Setting](db).Where("key IN ?", keys).Find(m.ctx)
	if err != nil {
		resp.Error = fmt.Sprintf("failed to list settings: %v", err)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}
	for _, s := range list {
		resp.KVMap[s.Key] = s.Value
	}
	return
}

type RespSaveSetting struct {
	Error string `json:"error"`
}

func (m *Model) SaveSetting(key string, value string) (resp RespSaveSetting) {
	if len(key) == 0 && len(value) == 0 {
		resp.Error = fmt.Sprintf("invalid setting, key: %s, value: %s", key, value)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}
	tx := db.Save(&Setting{key, value})
	if tx.Error != nil {
		resp.Error = fmt.Sprintf("failed to save setting, key: %s, value: %s, error: %v", key, value, tx.Error)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}
	return
}

// #endregion

// #region Tools

type ToolBrief struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func (t *ToolBrief) TableName() string {
	return "tools"
}

type RespGetToolList struct {
	Error string      `json:"error"`
	List  []ToolBrief `json:"list"`
}

func (m *Model) GetToolList() (resp RespGetToolList) {
	list, err := gorm.G[ToolBrief](db).Select("id", "name", "description").Order("name").Find(m.ctx)
	if err != nil {
		resp.Error = fmt.Sprintf("failed to list tools: %v", err)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}
	resp.List = list
	return
}

type RespGetTool struct {
	Error string `json:"error"`
	Item  Tool   `json:"item"`
}

func (m *Model) GetTool(id int) (resp RespGetTool) {
	var err error
	resp.Item, err = gorm.G[Tool](db).Where("id = ?", id).Take(m.ctx)
	if err != nil {
		resp.Error = fmt.Sprintf("failed to get tool detail: %v", err)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}
	return
}

type RespGetCommandLineTool struct {
	Error string          `json:"error"`
	Item  CommandLineTool `json:"item"`
}

func (m *Model) GetCommandLineTool(id int) (resp RespGetCommandLineTool) {
	// Get base tool from database
	tool, err := gorm.G[Tool](db).Where("id = ?", id).Take(m.ctx)
	if err != nil {
		resp.Error = fmt.Sprintf("failed to get tool: %v", err)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}

	// Set base tool info
	// resp.Item.Tool = tool

	// Evaluate the tool with default parameters using frontend WebWorker
	if tool.Code != "" && tool.DefaultParams != "" {
		toolData, err := EvalTool(m.ctx, tool.Code, tool.DefaultParams)
		if err != nil {
			resp.Error = fmt.Sprintf("failed to evaluate tool: %v", err)
			if m.ctx != nil {
				runtime.LogError(m.ctx, resp.Error)
			}
			return
		}
		// Parse the evaluated command line tool
		if err := json.Unmarshal(toolData, &resp.Item); err != nil {
			resp.Error = fmt.Sprintf("failed to parse command line tool: %v", err)
			if m.ctx != nil {
				runtime.LogError(m.ctx, resp.Error)
			}
			return
		}
		resp.Item.BaseModel = tool.BaseModel
	}

	return
}

type RespGetHTTPTool struct {
	Error string   `json:"error"`
	Item  HTTPTool `json:"item"`
}

func (m *Model) GetHTTPTool(id int) (resp RespGetHTTPTool) {
	// Get base tool from database
	tool, err := gorm.G[Tool](db).Where("id = ?", id).Take(m.ctx)
	if err != nil {
		resp.Error = fmt.Sprintf("failed to get tool: %v", err)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}

	// Set base tool info
	// resp.Item.Tool = tool

	// Evaluate the tool with default parameters using frontend WebWorker
	if tool.Code != "" && tool.DefaultParams != "" {
		toolData, err := EvalTool(m.ctx, tool.Code, tool.DefaultParams)
		if err != nil {
			resp.Error = fmt.Sprintf("failed to evaluate tool: %v", err)
			if m.ctx != nil {
				runtime.LogError(m.ctx, resp.Error)
			}
			return
		}

		// Parse the evaluated HTTP tool
		if err := json.Unmarshal(toolData, &resp.Item); err != nil {
			resp.Error = fmt.Sprintf("failed to parse HTTP tool: %v", err)
			if m.ctx != nil {
				runtime.LogError(m.ctx, resp.Error)
			}
			return
		}
		resp.Item.BaseModel = tool.BaseModel
	}

	return
}

type RespGetToolTestcaseList struct {
	Error string         `json:"error"`
	List  []ToolTestcase `json:"list"`
}

func (m *Model) GetToolTestcaseList(toolName string) (resp RespGetToolTestcaseList) {
	list, err := gorm.G[ToolTestcase](db).Where("tool_name = ?", toolName).Find(m.ctx)
	if err != nil {
		resp.Error = fmt.Sprintf("failed to list tool testcases: %v", err)
		if m.ctx != nil {
			runtime.LogError(m.ctx, resp.Error)
		}
		return
	}
	resp.List = list
	return
}

// #endregion
