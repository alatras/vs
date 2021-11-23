package middleware

import (
	"context"
	"net/http"
	"validation-service/config"
	"validation-service/enums/contextKey"

	"github.com/google/uuid"
)

func SetContextWithHeaders(l *config.Log) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			var traceId string
			var correlationId string

			if headerTraceId := r.Header.Get(l.TraceIdHeader); headerTraceId != "" {
				traceId = headerTraceId
			} else {
				traceId = uuid.New().String()
			}

			if headerCorrelationId := r.Header.Get("correlation_id"); headerCorrelationId != "" {
				correlationId = headerCorrelationId
			} else {
				correlationId = uuid.New().String()
			}

			ctx = context.WithValue(ctx, contextKey.TraceId, traceId)
			ctx = context.WithValue(ctx, contextKey.CorrelationId, correlationId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
