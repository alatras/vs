package cmd

import (
	"bitbucket.verifone.com/validation-service/app/validateTransaction"
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"bitbucket.verifone.com/validation-service/transaction"
	"context"
	"encoding/json"
	"github.com/google/uuid"
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

	var traceId string

	if headerTraceId := r.Header.Get("X-TRACE-ID"); headerTraceId != "" {
		traceId = headerTraceId
	} else {
		traceId = uuid.New().String()
	}

	ctx = context.WithValue(ctx, contextKey.TraceId, traceId)

	reportChan, errChan := h.validator.Enqueue(ctx, trx)

	select {
	case report := <-reportChan:
		reportJson, marshalErr := json.Marshal(report)

		if marshalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(reportJson)
	case err := <-errChan:
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
