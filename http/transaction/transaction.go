package transaction

import (
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
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

	r.Post("/validate", rs.Validate)

	return r
}
