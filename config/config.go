package config

import (
	"errors"
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

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverAddress, err := getEnv("SERVER_ADRESS")
	if err != nil {
		return nil, err
	}

	dbPath, err := getEnv("DATABASE_PATH")
	if err != nil {
		return nil, err
	}

	secretKey, err := getEnv("SECRET_KEY")
	if err != nil {
		return nil, err
	}

	log.Printf("Loaded .env file")

	return &Config{
		ServerAddress: serverAddress,
		DBPath:        dbPath,
		SecretKey:     secretKey,
	}, nil
}

func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", errors.New("environment variable not set: " + key)
	}

	return value, nil
}
