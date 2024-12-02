package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv(key)
}

func IsDev() bool {
	return GetEnv("ENVIRONMENT") == "development"
}
