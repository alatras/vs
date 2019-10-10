package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/app/listRuleSet"
	"bitbucket.verifone.com/validation-service/app/updateRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
)

type Resource struct {
	logger                  *logger.Logger
	createRuleSetAppFactory func() createRuleSet.CreateRuleSet
	listRuleSetAppFactory   func() listRuleSet.ListRuleSet
	getRuleSetAppFactory    func() getRuleSet.GetRuleSet
	deleteRuleSetAppFactory func() deleteRuleSet.DeleteRuleSet
	updateRuleSetAppFactory func() updateRuleSet.UpdateRuleSet
}

func NewResource(
	logger *logger.Logger,
	createRuleSetAppFactory func() createRuleSet.CreateRuleSet,
	getRuleSetAppFactory func() getRuleSet.GetRuleSet,
	deleteRuleSetAppFactory func() deleteRuleSet.DeleteRuleSet,
	listRuleSetAppFactory func() listRuleSet.ListRuleSet,
	updateRuleSetAppFactory func() updateRuleSet.UpdateRuleSet,
) Resource {
	return Resource{
		logger:                  logger,
		createRuleSetAppFactory: createRuleSetAppFactory,
		listRuleSetAppFactory:   listRuleSetAppFactory,
		getRuleSetAppFactory:    getRuleSetAppFactory,
		deleteRuleSetAppFactory: deleteRuleSetAppFactory,
		updateRuleSetAppFactory: updateRuleSetAppFactory,
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
