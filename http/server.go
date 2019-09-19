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
	router chi.Router
}

func NewServer(port string, r chi.Router) *Server {
	return &Server{
		port: port,
		router: r,
	}
}

func (s *Server) Start() error {
	r := s.router

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