package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Request defines the expected JSON structure for the input.
type Request struct {
	Name string `json:"name"`
}

// Response defines the expected JSON structure for the output.
type Response struct {
	Message string `json:"message"`
}

var server *http.Server

// HelloHandler handles POST requests to /hello.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request: Invalid JSON format", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Bad Request: Name is required", http.StatusBadRequest)
		return
	}

	resp := Response{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// TerminateHandler handles GET requests to /terminate and initiates graceful shutdown.
func TerminateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Server is shutting down...")
	fmt.Println("Shutdown signal received via /terminate")

	// Trigger shutdown in a separate goroutine to allow the response to be sent
	go ShutdownServer()
}

// ShutdownServer gracefully shuts down the server.
func ShutdownServer() {
	if server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Server Shutdown Failed:%+v\n", err)
		}
		fmt.Println("Server gracefully stopped")
	}
}

// StartServer starts the server on port 8080.
func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/terminate", TerminateHandler)

	server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server starting on port 8080...")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("HTTP server ListenAndServe: %v", err)
	}
}
