package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(addr string) {
	var err error

	db, err = sql.Open("sqlite3", addr)

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
