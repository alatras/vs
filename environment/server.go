package environment

import (
	"log"

	"github.com/joho/godotenv"
)

// Read .env if exists
func Read() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Print("No .env file")
	}
}
