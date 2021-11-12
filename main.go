package main

import (
	"log"
	"os"
	"validation-service/cmd"
	"validation-service/config"
)

func main() {
	config.Read()

	// log.Println("::::: config.App.AppD.AppName :::", config.App.AppD.AppName)
	// log.Println("::::: config.App.Log.Format :::", config.App.Log.Format)
	// log.Println("::::: config.App.HTTPPort :::", config.App.HTTPPort)
	// log.Println("::::: config.App.HTTPPort :::", config.App.Mongo.RetryMilliseconds)

	err := cmd.StartServer(config.App)

	if err != nil {
		log.Printf("[ERROR] server command failed with error: %+v", err)
		os.Exit(1)
	}
}
