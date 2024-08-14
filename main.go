package main

import (
	"log"
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/config"
	"github.com/hilbertgreveling/dnd-character-api/db"
	"github.com/hilbertgreveling/dnd-character-api/middleware"
	"github.com/hilbertgreveling/dnd-character-api/routes"
)

func main() {
	cfg := config.LoadConfig()

	db.InitDB()
	defer db.CloseDB()

	router := routes.SetupRoutes()

	stack := middleware.CreateStack(
		middleware.AuthMiddleware,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    cfg.ServerAddress,
		Handler: stack(router),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	log.Printf("Server running on %s", cfg.ServerAddress)
}
