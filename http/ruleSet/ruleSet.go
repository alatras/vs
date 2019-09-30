package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
	"net/http"
)

type Resource struct {
	logger                  *logger.Logger
	createRulesetAppFactory func() createRuleSet.CreateRuleSet
	listRulesetAppFactory   func() listRuleSet.ListRuleSet
	getRuleSetAppFactory    func() getRuleSet.GetRuleSet
	deleteRuleSetAppFactory func() deleteRuleSet.DeleteRuleSet
}

func NewResource(
	logger *logger.Logger,
	createRulesetAppFactory func() createRuleSet.CreateRuleSet,
	getRuleSetAppFactory func() getRuleSet.GetRuleSet,
	deleteRuleSetAppFactory func() deleteRuleSet.DeleteRuleSet,
	listRulesetAppFactory func() listRuleSet.ListRuleSet,
) Resource {
	return Resource{
		logger:                  logger,
		createRulesetAppFactory: createRulesetAppFactory,
		listRulesetAppFactory:   listRulesetAppFactory,
		getRuleSetAppFactory:    getRuleSetAppFactory,
		deleteRuleSetAppFactory: deleteRuleSetAppFactory,
	}
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/{id}/rulesets", rs.List)
	r.Post("/{id}/rulesets", rs.Create)

	r.Route("/{id}/rulesets/{ruleSetId}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
	})

	return r
}

func (rs Resource) Update(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Update rule sets"))
}
