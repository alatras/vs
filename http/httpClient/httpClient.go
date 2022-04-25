package httpClient

import (
	"validation-service/config"
	"validation-service/logger"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	logger          *logger.Logger
	config          *config.Server
	instrumentation *instrumentation
	record          *logger.LogRecord
	restyClient     *resty.Client
}

func NewHttpClient(
	logger *logger.Logger,
	config *config.Server,
	record *logger.LogRecord,
	restyClient *resty.Client,
) HttpClient {
	return HttpClient{
		instrumentation: newInstrumentation(logger, record),
		config:          config,
		logger:          logger,
		record:          record,
		restyClient:     restyClient,
	}
}
