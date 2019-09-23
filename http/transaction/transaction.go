package transaction

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

// Routes creates a REST router for transaction resources
func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/validate", rs.Test)
	r.Post("/validate", rs.Validate)

	return r
}

func (rs Resource) Test(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("test route added"))
}
