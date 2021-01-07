package http

import (
	"fmt"
	"net/http"
	"time"
	"validation-service/app/createRuleSet"
	"validation-service/app/deleteRuleSet"
	"validation-service/app/getRuleSet"
	"validation-service/app/listAncestorsRuleSet"
	"validation-service/app/listDescendantsRuleSet"
	"validation-service/app/listRuleSet"
	"validation-service/app/updateRuleSet"
	"validation-service/app/validateTransaction"
	"validation-service/http/healthCheck"
	httpMiddleware "validation-service/http/middleware"
	httpRuleSet "validation-service/http/ruleSet"
	"validation-service/http/transaction"
	"validation-service/logger"
	"validation-service/ruleSet"

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
