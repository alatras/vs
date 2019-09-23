package http

import (
	"bitbucket.verifone.com/validation-service/http/healthCheck"
	"bitbucket.verifone.com/validation-service/http/ruleset"
	"bitbucket.verifone.com/validation-service/http/transaction"
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

type Server struct {
	port   string
	router chi.Router
	logger *logger.Logger
}

func NewServer(port string, r chi.Router, l *logger.Logger) *Server {
	return &Server{
		port:   port,
		router: r,
		logger: l,
	}
}

func (s *Server) Start() error {
	r := s.router

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/healthCheck", healthCheck.NewResource(s.logger).Routes())
	r.Mount("/", transaction.NewResource(s.logger).Routes())
	r.Mount("/entities", ruleset.NewResource(s.logger).Routes())

	err := http.ListenAndServe(s.port, r)

	if err != nil {
		return err
	}

	return nil
}
