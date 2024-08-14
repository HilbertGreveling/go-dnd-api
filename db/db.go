package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/hilbertgreveling/dnd-character-api/config"
)

var db *sql.DB

func InitDB() {
	var err error
	cfg := config.GetConfig()
	db, err = sql.Open("sqlite3", cfg.DatabasePath)

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	log.Printf("Database OK!")
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database not initialized")
	}

	return db
}

func CloseDB() error {
	return db.Close()
}
