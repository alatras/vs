package healthCheck

import (
	"github.com/go-chi/chi"
	"net/http"
)

type Resource struct {}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.HealthCheck)

	return r
}

func (rs Resource) HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ =w.Write([]byte("Service is Up"))
}