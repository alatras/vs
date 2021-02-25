package environment

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

// Config .env if exists
func Config() {
	err := godotenv.Load(".env")

	if err != nil && os.Getenv("MONGO_URL") == "" {
		log.Print("Pass environment in 'run' command if deploying.")
	}
}

// Get env var or default
func Get(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		log.Print("[WARN] key" + key + " is not in environment. Returning fallback.")
		return fallback
	}

	return value
}

// GetDigits gets digits only of env var
func GetDigits(key string, fallback int) int {
	defer func() {
		if recover() != nil {
			log.Print("[WARN] failed to get digits only of for env var " + key)
		}
	}()

	value := os.Getenv(key)
	if len(value) == 0 {
		log.Print("[WARN] key" + key + " is not in environment. Returning fallback.")
		return fallback
	}

	d := regexp.MustCompile("[0-9]+").FindAllString(value, -1)[0]
	i, err := strconv.Atoi(d)
	if err != nil {
		log.Print("[WARN] key" + key + " is not in environment. Returning fallback.")
		return fallback
	}

	return i
}
