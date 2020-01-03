package transaction

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/logger"
	"github.com/go-chi/chi"
)

type Resource struct {
	logger *logger.Logger
	app    validateTransaction.ValidatorService
}

func NewResource(l *logger.Logger, a validateTransaction.ValidatorService) Resource {
	return Resource{
		logger: l,
		app:    a,
	}
}

// Routes creates a REST router for transaction resources
func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/validate", rs.Validate)

	return r
}
