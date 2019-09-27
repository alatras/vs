package http

import (
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/http/healthCheck"
	httpMiddleware "bitbucket.verifone.com/validation-service/http/middleware"
	"bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/http/transaction"
	"bitbucket.verifone.com/validation-service/logger"
	"fmt"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

type Server struct {
	port                   int
	router                 chi.Router
	logger                 *logger.Logger
	validateTransactionApp *validateTransaction.ValidatorService
	getRuleSetAppFactory   func() getRuleSet.GetRuleSet
}

func NewServer(
	port int,
	r chi.Router,
	l *logger.Logger,
	v *validateTransaction.ValidatorService,
	getRuleSetAppFactory func() getRuleSet.GetRuleSet,
) *Server {
	return &Server{
		port:                   port,
		router:                 r,
		logger:                 l,
		validateTransactionApp: v,
		getRuleSetAppFactory:   getRuleSetAppFactory,
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
	r.Mount("/entities", ruleSet.NewResource(s.logger, s.getRuleSetAppFactory).Routes())

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)

	if err != nil {
		e := fmt.Errorf("failed to the router %v", err)
		return e
	}

	return nil
}
