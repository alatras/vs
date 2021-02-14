package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Read .env if exists
func Read() {
	err := godotenv.Load(".env")

	if err != nil && os.Getenv("MONGO_URL") == "" {
		log.Print("Pass environment in 'run' command if deploying.")
	}
}
