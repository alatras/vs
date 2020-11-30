package http

import (
	"fmt"
	"net/http"
	"time"

	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/app/listDescendantsRuleSet"
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/http/healthCheck"
	httpMiddleware "bitbucket.verifone.com/validation-service/http/middleware"
	httpRuleSet "bitbucket.verifone.com/validation-service/http/ruleSet"
	"bitbucket.verifone.com/validation-service/http/transaction"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Server struct {
	port                             int
	router                           chi.Router
	logger                           *logger.Logger
	ruleSetRepository                ruleSet.Repository
	validateTransactionService       validateTransaction.ValidatorService
	createRuleSetAppFactory          func() createRuleSet.CreateRuleSet
	listRuleSetAppFactory            func() listRuleSet.ListRuleSet
	listAncestorsRuleSetAppFactory   func() listAncestorsRuleSet.ListAncestorsRuleSet
	listDescendantsRuleSetAppFactory func() listDescendantsRuleSet.ListDescendantsRuleSet
	getRuleSetAppFactory             func() getRuleSet.GetRuleSet
	deleteRuleSetAppFactory          func() deleteRuleSet.DeleteRuleSet
	updateRuleSetAppFactory          func() updateRuleSet.UpdateRuleSet
}

func NewServer(
	port int,
	router chi.Router,
	logger *logger.Logger,
	ruleSetRepository ruleSet.Repository,
	validateTransactionService validateTransaction.ValidatorService,
	createRuleSetAppFactory func() createRuleSet.CreateRuleSet,
	listRuleSetAppFactory func() listRuleSet.ListRuleSet,
	listAncestorsRuleSetAppFactory func() listAncestorsRuleSet.ListAncestorsRuleSet,
	listDescendantsRuleSetAppFactory func() listDescendantsRuleSet.ListDescendantsRuleSet,
	getRuleSetAppFactory func() getRuleSet.GetRuleSet,
	deleteRuleSetAppFactory func() deleteRuleSet.DeleteRuleSet,
	updateRuleSetAppFactory func() updateRuleSet.UpdateRuleSet,
) *Server {
	return &Server{
		port:                             port,
		router:                           router,
		logger:                           logger,
		ruleSetRepository:                ruleSetRepository,
		validateTransactionService:       validateTransactionService,
		createRuleSetAppFactory:          createRuleSetAppFactory,
		listRuleSetAppFactory:            listRuleSetAppFactory,
		listAncestorsRuleSetAppFactory:   listAncestorsRuleSetAppFactory,
		listDescendantsRuleSetAppFactory: listDescendantsRuleSetAppFactory,
		getRuleSetAppFactory:             getRuleSetAppFactory,
		deleteRuleSetAppFactory:          deleteRuleSetAppFactory,
		updateRuleSetAppFactory:          updateRuleSetAppFactory,
	}
}

func (s *Server) Start() error {
	r := s.router

	r.Use(httpMiddleware.SetContextWithTraceId)
	r.Use(httpMiddleware.Logger(s.logger))
	r.Use(chiMiddleware.URLFormat)
	r.Use(httpMiddleware.SetContextWithBusinessTransaction)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/healthCheck", healthCheck.NewResource(s.logger, s.ruleSetRepository).Routes())
	r.Mount("/transaction", transaction.NewResource(s.logger, s.validateTransactionService).Routes())
	r.Mount(
		"/entities",
		httpRuleSet.NewResource(
			s.logger,
			s.createRuleSetAppFactory,
			s.getRuleSetAppFactory,
			s.deleteRuleSetAppFactory,
			s.listRuleSetAppFactory,
			s.listAncestorsRuleSetAppFactory,
			s.listDescendantsRuleSetAppFactory,
			s.updateRuleSetAppFactory,
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
