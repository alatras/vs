package middleware

import (
	appd "bitbucket.verifone.com/validation-service/appdynamics"
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"context"
	"github.com/google/uuid"
	"net/http"
)

func SetContextWithBusinessTransaction(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		correlationHeader := r.Header.Get(appd.APPD_CORRELATION_HEADER_NAME)

		businessTransaction := appd.StartBT(r.URL.Path, correlationHeader)

		transactionUid := uuid.New().String()
		appd.StoreBT(businessTransaction, transactionUid)

		ctx = context.WithValue(ctx, contextKey.BusinessTransaction, transactionUid)
		next.ServeHTTP(w, r.WithContext(ctx))

		appd.EndBT(businessTransaction)
	})
}
