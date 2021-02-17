package listDescendantsRuleSet

import (
	"context"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
)

type metadata = logger.Metadata
type field = logger.Field

type instrumentation struct {
	logger    *logger.Logger
	startedAt time.Time
}

var fileLog = logger.FileLogger{}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	scope := "ListDescendantsRuleSet"
	fileLog.Scope = scope

	return &instrumentation{
		logger: logger.Scoped(scope),
	}
}

func (i *instrumentation) setContext(ctx context.Context) {
	if traceId, ok := ctx.Value(contextKey.TraceId).(string); ok {
		i.logger = i.logger.WithTraceId(traceId)
		fileLog.TraceID = traceId
	}
}

func (i *instrumentation) setMetadata(metadata metadata) {
	i.logger = i.logger.WithMetadata(metadata)
	fileLog.Metadata = metadata
}

func (i *instrumentation) startListingDescendantsRuleSet() {
	entry := i.logger.Output
	msg := "starting listing descendants rule sets"

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) finishListingDescendantsRuleSet() {
	duration := time.Since(i.startedAt)
	entry := i.logger.Output.WithField("duration", duration)
	msg := "finished listing descendants rule set"

	fileLog.Message = msg
	fileLog.Fields = []field{{"duration", duration}}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) failedListingDescendantsRuleSet(err error) {
	entry := i.logger.Output.Logger.WithError(err)
	msg := "error fetching descendants rule sets from db"

	fileLog.Message = msg
	fileLog.Error = err
	fileLog.Fields = []field{}

	entry.Error(msg)
	logger.WriteToFile(fileLog)
}
