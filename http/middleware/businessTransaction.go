package middleware

import (
	"context"
	"net/http"
	"regexp"
	appd "validation-service/appdynamics"
	"validation-service/enums/contextKey"

	"github.com/google/uuid"
)

func SetContextWithBusinessTransaction(next http.Handler) http.Handler {
	var uuidRegex = regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		path := uuidRegex.ReplaceAllString(r.URL.Path, `{uuid}`)

		transactionName := r.Method + " " + path

		correlationHeader := r.Header.Get(appd.APPD_CORRELATION_HEADER_NAME)

		businessTransaction := appd.StartBT(transactionName, correlationHeader)

		transactionUid := uuid.New().String()
		appd.StoreBT(businessTransaction, transactionUid)

		ctx = context.WithValue(ctx, contextKey.BusinessTransaction, transactionUid)
		next.ServeHTTP(w, r.WithContext(ctx))

		appd.EndBT(businessTransaction)
	})
}
