package hub

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// StartHub starts the HTTP server and initializes tool evaluation listener
func StartHub(ctx context.Context) {
	// Initialize the global event listener for tool evaluation
	InitToolEvalListener(ctx)

	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/registerTool", registerToolHandler(ctx))
	http.HandleFunc("/api/callTool", callToolHandler(ctx))
	// http.HandleFunc("/ws/callStreamTool", callStreamToolHandler)
	// http.HandleFunc("/terminal", createTerminalHandler(ctx))

	server := &http.Server{Addr: "0.0.0.0:9573"}
	go func() {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}()
	log.Fatal(server.ListenAndServe())
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func registerToolHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		registerTool(ctx, w, r)
	}
}

func callToolHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		callTool(ctx, w, r)
	}
}
