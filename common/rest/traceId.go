package rest

import (
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"context"
	"github.com/google/uuid"
	"net/http"
)

func GetContextWithTraceId(r *http.Request) context.Context {
	ctx := r.Context()

	var traceId string

	if headerTraceId := r.Header.Get("X-TRACE-ID"); headerTraceId != "" {
		traceId = headerTraceId
	} else {
		traceId = uuid.New().String()
	}

	ctx = context.WithValue(ctx, contextKey.TraceId, traceId)

	return ctx
}
