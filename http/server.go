package http

import (
	"bitbucket.verifone.com/validation-service/http/healthCheck"
	"bitbucket.verifone.com/validation-service/http/ruleset"
	"bitbucket.verifone.com/validation-service/http/transaction"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/healthCheck", healthCheck.Resource{}.Routes())
	r.Mount("/", transaction.Resource{}.Routes())
	r.Mount("/entities", ruleset.Resource{}.Routes())

	err := http.ListenAndServe(s.port, r)

	if err != nil {
		return err
	}

	return nil
}