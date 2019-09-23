package http

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/http/healthCheck"
	httpMiddleware "bitbucket.verifone.com/validation-service/http/middleware"
	"bitbucket.verifone.com/validation-service/http/ruleset"
	"bitbucket.verifone.com/validation-service/http/transaction"
	"bitbucket.verifone.com/validation-service/logger"
	"fmt"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

type Server struct {
	port                   string
	router                 chi.Router
	logger                 *logger.Logger
	validateTransactionApp *validateTransaction.ValidatorService
}

func NewServer(port string, r chi.Router, l *logger.Logger, v *validateTransaction.ValidatorService) *Server {
	return &Server{
		port:                   port,
		router:                 r,
		logger:                 l,
		validateTransactionApp: v,
	}
}

func (s *Server) Start() error {
	r := s.router

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.URLFormat)
	r.Use(httpMiddleware.SetContextWithTraceId)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/healthCheck", healthCheck.NewResource(s.logger).Routes())
	r.Mount("/", transaction.NewResource(s.logger, s.validateTransactionApp).Routes())
	r.Mount("/entities", ruleset.NewResource(s.logger).Routes())

	err := http.ListenAndServe(s.port, r)

	if err != nil {
		e := fmt.Errorf("failed to the router %v", err)
		return e
	}

	return nil
}
