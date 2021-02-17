package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config .env if exists
func Config() {
	err := godotenv.Load(".env")

	if err != nil && os.Getenv("MONGO_URL") == "" {
		log.Print("Pass environment in 'run' command if deploying.")
	}
}

// Get : get env var or default
func Get(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		log.Println("[WARN] key" + key + " is not in environment. Returning fallback.")
		return fallback
	}

	return value
}
