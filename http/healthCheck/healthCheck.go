package healthCheck

import (
	appd "bitbucket.verifone.com/validation-service/appdynamics"
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
	appDCorrelationHeader := r.Header.Get(appd.APPD_CORRELATION_HEADER_NAME)
	businessTransaction := appd.StartBT("Health check", appDCorrelationHeader)
	appd.SetBTURL(businessTransaction, r.URL.Path)
	defer appd.EndBT(businessTransaction)

	err := rs.ruleSetRepo.Ping(r.Context())

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), false)

		rs.logger.Error.WithError(err).Error("Health check failed. Mongo is down.......")

		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
