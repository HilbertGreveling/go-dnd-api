package server

import (
	"log"
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/middleware"
	"github.com/hilbertgreveling/dnd-character-api/routes"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Serve() {
	mux := http.NewServeMux()
	mux = routes.SetupRoutes(mux)

	stack := middleware.CreateStack(
		middleware.CORS,
		middleware.AuthMiddleware,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: stack(mux),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	log.Printf("Server running on %s", s.addr)
}
