package middleware

import (
	"fmt"
	"net/http"
	"time"

	"validation-service/enums/contextKey"
	"validation-service/logger"

	chi "github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// RequestLogFormatter log formatter for HTTP requests
type RequestLogFormatter struct {
	logger *logger.Logger
}

// RequestLogEntry HTTP request log entry
type RequestLogEntry struct {
	output  *logrus.Entry
	request *http.Request
	fmt     chi.LogFormatter
	msg     string
}

// Write called by chi when request has been processed
func (e *RequestLogEntry) Write(status, bytes int, elapsed time.Duration) {
	msg := fmt.Sprintf("%s - %d %s", e.msg, status, elapsed)

	output := e.output

	if bytes != 0 {
		output = output.WithField("response", fmt.Sprintf("%d bytes", bytes))
	}

	output.Info(msg)
}

// Panic called by chi when request caused a panic
func (e *RequestLogEntry) Panic(v interface{}, stack []byte) {
	panicEntry := e.fmt.NewLogEntry(e.request).(*RequestLogEntry)
	e.output.Errorf("Panic: %s", panicEntry.msg)
	e.output.Errorf("Stack: %s", string(stack))
}

// NewLogEntry called by chi when new request has arrived
func (f RequestLogFormatter) NewLogEntry(r *http.Request) chi.LogEntry {
	scheme := "http"

	if r.TLS != nil {
		scheme = "https"
	}

	msg := fmt.Sprintf("%s %s://%s%s %s", r.Method, scheme, r.Host, r.RequestURI, r.Proto)

	output := f.logger.Output.WithField("from", r.RemoteAddr)

	if r.ContentLength != 0 {
		output = output.WithField("body", fmt.Sprintf("%d bytes", r.ContentLength))
	}

	if traceID, ok := r.Context().Value(contextKey.TraceId).(string); ok {
		output = output.WithField("trace_id", traceID)
	}

	return &RequestLogEntry{
		output:  output,
		request: r,
		msg:     msg,
		fmt:     f,
	}
}

// Func function which is accepted as chi middleware
type Func = func(next http.Handler) http.Handler

// Logger creates logger middleware which will output to specified logger
func Logger(logger *logger.Logger) Func {
	formatter := RequestLogFormatter{logger}

	return func(next http.Handler) http.Handler {
		requestLogger := chi.RequestLogger(&formatter)
		return requestLogger(next)
	}
}
