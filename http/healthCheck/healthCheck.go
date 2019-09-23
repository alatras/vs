package healthCheck

import (
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
	"net/http"
)

type Resource struct {
	logger *logger.Logger
}

func NewResource(l *logger.Logger) Resource {
	return Resource{
		logger: l,
	}
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.HealthCheck)

	return r
}

func (rs Resource) HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(""))

	if err != nil {
		rs.logger.Error.WithError(err).Error("Health check failed")
	}
}
