package main

import (
	"github.com/hilbertgreveling/dnd-character-api/config"
	"github.com/hilbertgreveling/dnd-character-api/db"
	"github.com/hilbertgreveling/dnd-character-api/server"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg.DatabasePath)
	defer db.CloseDB()

	svr := server.NewAPIServer(cfg.ServerAddress)
	svr.Serve()
}
