package listRuleSet

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
	scope := "ListRuleSet"
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

func (i *instrumentation) startListingRuleSet() {
	entry := i.logger.Output
	now := time.Now()
	msg := "starting listing the rule sets"
	i.startedAt = now

	fileLog.Message = msg
	fileLog.StartedAt = now
	fileLog.Fields = []field{}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) finishListingRuleSet() {
	duration := time.Since(i.startedAt)
	entry := i.logger.Output.WithField("duration", duration)
	msg := "finished listing rule set"

	fileLog.Message = msg
	fileLog.Fields = []field{{"duration", duration}}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) failedListingRuleSet(err error) {
	entry := i.logger.Output.Logger.WithError(err)
	msg := "error fetching ruleset from db"

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.Error = err

	entry.Error(msg)
	logger.WriteToFile(fileLog)
}
