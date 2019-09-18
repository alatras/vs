package cmd

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"bitbucket.verifone.com/validation-service/transaction"
	"context"
	"encoding/json"
	"net/http"
)

type HttpServer struct {
	port      string
	validator validateTransaction.ValidatorService
}

func NewHttpServer(port string, validator validateTransaction.ValidatorService) *HttpServer {
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

	ctx := r.Context()

	if traceId := r.Header.Get("X-TRACE-ID"); traceId != "" {
		ctx = context.WithValue(ctx, contextKey.TraceId, traceId)
	}

	report := <-h.validator.Enqueue(trx, ctx)
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
