package app

import (
	"CodeXecutor/internal/app/handler"
	"log"
	"net/http"
)

// Server represents the application server.
type Server struct {
	// Add server-related fields and dependencies here
}

// NewServer initializes and returns a new Server instance.
func NewServer() *Server {
	return &Server{}
}

// Start starts the application server.
func (s *Server) Start() {
	// Initialize routes and start the HTTP server
	http.HandleFunc("/", handler.HandleCodeSubmission)

	go func() {
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
}

// Stop gracefully stops the application server.
func (s *Server) Stop() {
	// Implement graceful shutdown logic here
}