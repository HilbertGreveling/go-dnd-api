package server

import (
	"log"
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/handlers"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Serve() {
	server := http.Server{
		Addr:    s.addr,
		Handler: handlers.SetupHandlers(),
	}

	log.Printf("Server running on %s", s.addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
