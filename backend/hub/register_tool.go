package hub

import (
	"context"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type BodyRegisterTool struct {
	Tool Tool `json:"tool"`
}

func registerTool(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithoutCancel(ctx)
	var body BodyRegisterTool
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tool := body.Tool
	oTool, err := gorm.G[Tool](db).Where("name = ?", tool.Name).Take(ctx)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		// new tool
		err = gorm.G[Tool](db).Create(ctx, &tool)
		if err != nil {
			http.Error(w, "Failed to create tool", http.StatusInternalServerError)
			return
		}
		return
	}
	tool.ID = oTool.ID
	tool.CreatedAt = oTool.CreatedAt
	tool.UpdatedAt = oTool.UpdatedAt
	_, err = gorm.G[Tool](db).Updates(ctx, tool)
	if err != nil {
		http.Error(w, "Failed to update tool", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message": "register tool done"}`))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
