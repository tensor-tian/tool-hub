package hub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"

	"tool-hub/backend/hub/cmd"
)

// BodyCallTool represents the request body for calling a tool
type BodyCallTool struct {
	Name       string `json:"name"`
	Parameters string `json:"parameters"`
}

func callTool(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithoutCancel(ctx)
	var body BodyCallTool
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch tool from database
	tool, err := gorm.G[Tool](db).Where("name = ?", body.Name).Take(ctx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, fmt.Sprintf("Tool not found: %s", body.Name), http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Evaluate tool using frontend WebWorker
	toolData, err := EvalTool(ctx, tool.Code, body.Parameters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the tool based on category
	var commandLineTool CommandLineTool
	if err := json.Unmarshal(toolData, &commandLineTool); err != nil {
		http.Error(w, "Failed to parse tool response", http.StatusInternalServerError)
		return
	}

	// Parse environment variables
	var envMap map[string]string
	if commandLineTool.Extra.Env != "" {
		if err := json.Unmarshal([]byte(commandLineTool.Extra.Env), &envMap); err != nil {
			envMap = nil
		}
	}

	// Parse timeout from tool configuration
	var timeout time.Duration
	if commandLineTool.Timeout != "" {
		timeout, err = time.ParseDuration(commandLineTool.Timeout)
		if err != nil {
			timeout = 0 // Use default (no timeout)
		}
	}

	// Execute the command using shared runner
	input := cmd.Input{
		Reader: bytes.NewReader([]byte(commandLineTool.Extra.Stdin)),
		Options: cmd.StreamOptions{
			Cwd:     commandLineTool.Extra.WD,
			Env:     envMap,
			Shell:   commandLineTool.Extra.Sh,
			Timeout: timeout,
		},
		Command: []string{commandLineTool.Extra.Cmd},
	}

	out, err := cmd.SharedRunner.Run(input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Command execution failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
