package instrumentation

import (
	"bitbucket.verifone.com/validation-service/infra/logger"
)

type MainInstrumentation struct {
	logger *logger.Logger
}

func NewMainInstrumentation(logger *logger.Logger) *MainInstrumentation {
	return &MainInstrumentation{
		logger: logger.Scoped("Main"),
	}
}

func (i *MainInstrumentation) SetTraceId(traceId string) {
	i.logger = i.logger.WithTraceId(traceId)
}

func (i *MainInstrumentation) SetMetadata(metadata interface{}) {
	i.logger = i.logger.WithMetadata(metadata)
}

func (i *MainInstrumentation) FailedToInitRuleSetRepository(error error) {
	i.logger.Error.WithError(error).Error("Failed to initialize RuleSetRepository")
}

func (i *MainInstrumentation) StartingRestApiServer(port int) {
	i.logger.Output.Infof("Starting REST API server at port %d", port)
}

func (i *MainInstrumentation) FailedToStartRestApiServer(error error) {
	i.logger.Error.WithError(error).Error("Failed to start REST API server")
}

