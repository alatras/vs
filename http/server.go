package http

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/http/healthCheck"
	httpMiddleware "bitbucket.verifone.com/validation-service/http/middleware"
	httpRuleSet "bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/http/transaction"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"fmt"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type Server struct {
	port                       int
	router                     chi.Router
	logger                     *logger.Logger
	ruleSetRepository          ruleSet.Repository
	validateTransactionService *validateTransaction.ValidatorService
	createRuleSetAppFactory    func() createRuleSet.CreateRuleSet
	listRuleSetAppFactory      func() listRuleSet.ListRuleSet
	getRuleSetAppFactory       func() getRuleSet.GetRuleSet
	deleteRuleSetAppFactory    func() deleteRuleSet.DeleteRuleSet
}

func NewServer(
	port int,
	router chi.Router,
	logger *logger.Logger,
	ruleSetRepository ruleSet.Repository,
	validateTransactionService *validateTransaction.ValidatorService,
	createRuleSetAppFactory func() createRuleSet.CreateRuleSet,
	listRuleSetAppFactory func() listRuleSet.ListRuleSet,
	getRuleSetAppFactory func() getRuleSet.GetRuleSet,
	deleteRuleSetAppFactory func() deleteRuleSet.DeleteRuleSet,
) *Server {
	return &Server{
		port:                       port,
		router:                     router,
		logger:                     logger,
		ruleSetRepository:          ruleSetRepository,
		validateTransactionService: validateTransactionService,
		createRuleSetAppFactory:    createRuleSetAppFactory,
		listRuleSetAppFactory:      listRuleSetAppFactory,
		getRuleSetAppFactory:       getRuleSetAppFactory,
		deleteRuleSetAppFactory:    deleteRuleSetAppFactory,
	}
}

func (s *Server) Start() error {
	r := s.router

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.URLFormat)
	r.Use(httpMiddleware.SetContextWithTraceId)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/healthCheck", healthCheck.NewResource(s.logger).Routes())
	r.Mount("/", transaction.NewResource(s.logger, s.validateTransactionService).Routes())
	r.Mount(
		"/entities",
		httpRuleSet.NewResource(
			s.logger,
			s.createRuleSetAppFactory,
			s.getRuleSetAppFactory,
			s.deleteRuleSetAppFactory,
			s.listRuleSetAppFactory,
		).Routes(),
	)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	err := httpServer.ListenAndServe()

	if err != nil {
		e := fmt.Errorf("failed to the router %v", err)
		return e
	}

	return nil
}
