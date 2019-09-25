package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
	"net/http"
)

type Resource struct {
	logger                  *logger.Logger
	createRulesetAppFactory func() createRuleSet.CreateRuleset
}

func NewResource(
	logger *logger.Logger,
	createRulesetAppFactory func() createRuleSet.CreateRuleset,
) Resource {
	return Resource{
		logger:                  logger,
		createRulesetAppFactory: createRulesetAppFactory,
	}
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/{id}/rulesets", rs.List)
	r.Post("/{id}/rulesets", rs.Create)

	r.Route("/{id}/rulesets/{rulesetId}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
	})

	return r
}

func (rs Resource) List(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("List rule sets"))
}

func (rs Resource) Get(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Get rule sets"))
}

func (rs Resource) Update(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Update rule sets"))
}

func (rs Resource) Delete(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Delete rule sets"))
}
