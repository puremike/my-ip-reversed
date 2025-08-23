package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") == "development" {
		if err := godotenv.Load(); err != nil {
			log.Print("loading .env file failed")
		}
	}
}

func GetEnvString(key, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
