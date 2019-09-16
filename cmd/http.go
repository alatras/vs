package cmd

import (
	"bitbucket.verifone.com/validation-service/app"
	"bitbucket.verifone.com/validation-service/domain/transaction"
	"encoding/json"
	"net/http"
)

type HttpServer struct {
	port      string
	validator app.ValidatorService
}

func NewHttpServer(port string, validator app.ValidatorService) *HttpServer {
	return &HttpServer{
		port:      port,
		validator: validator,
	}
}

func (h *HttpServer) Start() error {
	http.HandleFunc("/validate", h.validateHandler)
	err := http.ListenAndServe(h.port, nil)

	if err != nil {
		return err
	}

	return nil
}

func (h *HttpServer) validateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	trx := transaction.Transaction{}
	err := json.NewDecoder(r.Body).Decode(&trx)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Bad Request"))
		return
	}

	report := <-h.validator.Enqueue(trx)
	reportJson, marshalErr := json.Marshal(report)

	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(reportJson)
}
