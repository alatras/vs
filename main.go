package main

import (
	"log"
	"os"
	"validation-service/cmd"
	"validation-service/config"
)

func main() {
	config.Read("")

	err := cmd.StartServer(config.App)

	if err != nil {
		log.Printf("[ERROR] server command failed with error: %+v", err)
		os.Exit(1)
	}
}
