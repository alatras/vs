package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/getRuleSet"
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
	"net/http"
)

type Resource struct {
	logger               *logger.Logger
	getRuleSetAppFactory func() getRuleSet.GetRuleSet
}

func NewResource(
	l *logger.Logger,
	getRuleSetAppFactory func() getRuleSet.GetRuleSet,
) Resource {
	return Resource{
		logger:               l,
		getRuleSetAppFactory: getRuleSetAppFactory,
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

func (rs Resource) List(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("List rule sets"))
}

func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Create rule sets"))
}

func (rs Resource) Update(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Update rule sets"))
}

func (rs Resource) Delete(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Delete rule sets"))
}
