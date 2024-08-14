package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DBPath        string
	SecretKey     string
}

var cfg *Config

func LoadConfig() *Config {
	if cfg == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}

		cfg = &Config{
			ServerAddress: getEnv("SERVER_ADDRESS", ""),
			DBPath:        getEnv("DB_PATH", ""),
			SecretKey:     getEnv("SECRET_KEY", ""),
		}

		if cfg.ServerAddress == "" || cfg.DBPath == "" || cfg.SecretKey == "" {
			log.Fatal("Missing required environment variables")
		}
	}

	log.Printf("Loaded .env file")

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
