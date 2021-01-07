package healthCheck

import (
	"net/http"
	appd "validation-service/appdynamics"
	"validation-service/enums/contextKey"
	"validation-service/logger"
	"validation-service/ruleSet"

	"github.com/go-chi/chi"
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
	ctx := r.Context()

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	err := rs.ruleSetRepo.Ping(ctx)

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)

		rs.logger.Error.WithError(err).Error("Health check failed. Mongo is down.......")

		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
