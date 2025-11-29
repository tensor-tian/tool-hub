package hub

import (
	"context"
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

type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type ToolDetail struct {
	Tool
	Parameters           string `json:"parameters"`
	LogLifeSpan          string `json:"logLifeSpan"`
	ConcurrencyGroupName string `json:"concurrencyGroupName"`
	Extra                string `json:"extra"`
}

type RespGetToolList struct {
	Error string `json:"error"`
	List  []Tool `json:"list"`
}

type RespToolDetail struct {
	Error string     `json:"error"`
	item  ToolDetail `json:"item"`
}

func (m *Model) GetTools() (resp RespGetToolList) {
	return
}

func (m *Model) GetToolDetail(name string) (resp RespToolDetail) {
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
