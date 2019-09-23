package middleware

import (
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"context"
	"github.com/google/uuid"
	"net/http"
)

func SetContextWithTraceId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var traceId string
		if headerTraceId := r.Header.Get("X-TRACE-ID"); headerTraceId != "" {
			traceId = headerTraceId
		} else {
			traceId = uuid.New().String()
		}

		ctx = context.WithValue(ctx, contextKey.TraceId, traceId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
