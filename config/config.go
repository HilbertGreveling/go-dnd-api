package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DatabasePath  string
	SecretKey     string
}

var cfg *Config

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Config: Error loading .env file")
	}

	log.Printf("Loaded .env file")

	cfg = &Config{
		ServerAddress: getEnv("SERVER_ADDRESS"),
		DatabasePath:  getEnv("DATABASE_PATH"),
		SecretKey:     getEnv("SECRET_KEY"),
	}

	log.Printf("Set env variables")

	return cfg
}

func GetConfig() *Config {
	if cfg == nil {
		log.Fatal("Config not initialized")
	}

	return cfg
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Config: Error loading env variable with key: ", key)
	}

	return value
}
