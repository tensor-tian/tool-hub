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

type Setting struct {
	Key   string `json:"key" gorm:"primarykey"`
	Value string `json:"value"`
}

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
