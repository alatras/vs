package middleware

import (
	"net/http"
	"strings"
	"validation-service/enums/contextKey"
	"validation-service/http/errorResponse"
	"validation-service/http/httpClient"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func CheckEntity(c *httpClient.HttpClient) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			var traceID string
			traceIdInterface := r.Context().Value(contextKey.TraceId)
			if traceIdInterface != nil {
				traceID = traceIdInterface.(string)
			}

			entityId := chi.URLParam(r, "id")

			token := r.Header.Get("Authorization")
			if token == "" {
				_ = render.Render(w, r, errorResponse.UnauthorizedRequest(entityId, "No auth token"))
				return
			}

			arr := strings.Split(token, "Bearer ")
			if len(arr) > 1 {
				token = strings.Split(token, "Bearer ")[1]
			}

			status := c.CheckEntity(token, entityId, traceID)
			if status != 200 {
				switch status {
				case 401:
					_ = render.Render(w, r, errorResponse.UnauthorizedRequest(entityId, errorResponse.EntityServiceUnauthorizeMessage))
				case 403:
					_ = render.Render(w, r, errorResponse.ForbiddenRequest(entityId, errorResponse.EntityServiceForbidMessage))
				case 404:
					_ = render.Render(w, r, errorResponse.NotFound(errorResponse.EntityServiceNotFoundMessage))
				case 400:
					_ = render.Render(w, r, errorResponse.MalformedParameters(errorResponse.EntityServiceMalformedEntiresMessage))
				case 500:
					_ = render.Render(w, r, errorResponse.UnexpectedError(errorResponse.EntityServiceUnexpectedErrorMessage))
				case 502:
					_ = render.Render(w, r, errorResponse.UnexpectedError(errorResponse.EntityServiceUnexpectedErrorMessage))
				case 512:
					_ = render.Render(w, r, errorResponse.UnexpectedError(errorResponse.UnexpectedErrorWithResty))
				default:
					_ = render.Render(w, r, errorResponse.HttpClientError(entityId, errorResponse.EntityServiceUnexpectedMessage))
				}
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
