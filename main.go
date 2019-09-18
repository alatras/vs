package main

import (
	//"bitbucket.verifone.com/validation-service/app"
	"bitbucket.verifone.com/validation-service/http"
	//"bitbucket.verifone.com/validation-service/infra/repository"
	"fmt"
	"log"
	"os"
)

var version = "unknown"

func main() {
	log.Printf("Validation Service %s\n", version)

	//ruleSetRepository, err := repository.NewStubRuleSetRepository()
	//
	//if err != nil {
	//	_, _ = fmt.Fprintln(os.Stderr, err)
	//	return
	//}
	//
	//validatorService := app.NewValidatorService(6, ruleSetRepository)


	err := http.NewServer(":8080").Start()

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
//package main
//
//import (
//	"errors"
//	"fmt"
//	"net/http"
//
//	"github.com/go-chi/chi"
//)
//
//type Handler func(w http.ResponseWriter, r *http.Request) error
//
//func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if err := h(w, r); err != nil {
//		// handle returned error here.
//		w.WriteHeader(503)
//		w.Write([]byte("bad"))
//	}
//}
//
//func main() {
//	r := chi.NewRouter()
//	r.Method("GET", "/", Handler(customHandler))
//	http.ListenAndServe(":8080", r)
//	fmt.Println("Server started successfully")
//}
//
//func customHandler(w http.ResponseWriter, r *http.Request) error {
//	q := r.URL.Query().Get("err")
//
//	if q != "" {
//		return errors.New(q)
//	}
//
//	w.Write([]byte("foo"))
//	return nil
//}
