package listAncestorsRuleSet

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
	scope := "ListAncestorsRuleSet"
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

func (i *instrumentation) startListingAncestorsRuleSet() {
	now := time.Now()
	entry := i.logger.Output
	msg := "starting listing ancestors rule sets"
	i.startedAt = now

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.StartedAt = now

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) finishListingAncestorsRuleSet() {
	since := time.Since(i.startedAt)
	entry := i.logger.Output.WithField("duration", since)
	msg := "finished listing ancestors rule set"

	fileLog.Message = msg
	fileLog.Fields = []field{{"duration", since}}

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) failedListingAncestorsRuleSet(err error) {
	entry := i.logger.Output.Logger.WithError(err)
	msg := "error fetching ancestors rule sets from db"

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.Error = err

	entry.Error(msg)
	logger.WriteToFile(fileLog)
}
