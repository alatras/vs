package healthCheck

import (
	"bitbucket.verifone.com/validation-service/entityService"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/go-chi/chi"
	"net/http"
)

type Resource struct {
	logger              *logger.Logger
	ruleSetRepo         ruleSet.Repository
	entityServiceClient entityService.EntityService
}

func NewResource(l *logger.Logger, r ruleSet.Repository, e entityService.EntityService) Resource {
	return Resource{
		logger:              l.Scoped("healthCheck"),
		ruleSetRepo:         r,
		entityServiceClient: e,
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

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rs.entityServiceClient.Ping()

	if err != nil {

		rs.logger.Error.WithError(err).Error("Health check failed. Entity Service is down.......")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
