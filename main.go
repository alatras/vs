package main

import (
	"bitbucket.verifone.com/validation-service/http"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"os"
)

var version = "unknown"

const port = ":8080"

func main() {
	log.Printf("Validation Service %s\n", version)

	err := http.NewServer(port, chi.NewRouter()).Start()

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
