package healthCheck

import (
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/go-chi/chi"
	"net/http"
)

type Resource struct {
	logger      *logger.Logger
	ruleSetRepo ruleSet.Repository
}

func NewResource(l *logger.Logger, r ruleSet.Repository) Resource {
	return Resource{
		logger:      l.Scoped("healthCheck"),
		ruleSetRepo: r,
	}
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.HealthCheck)

	return r
}

func (rs Resource) HealthCheck(w http.ResponseWriter, r *http.Request) {
	err := rs.ruleSetRepo.Ping(r.Context())

	if err != nil {

		rs.logger.Error.WithError(err).Error("Health check failed. Mongo is down.......")

		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
