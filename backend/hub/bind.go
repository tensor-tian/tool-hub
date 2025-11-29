package hub

import (
	"context"
	"os"
	"path"
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
