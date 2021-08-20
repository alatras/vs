package listDescendantsRuleSet

import (
	"context"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
)

type metadata = logger.Metadata

type instrumentation struct {
	logger    *logger.Logger
	startedAt time.Time
	record    *logger.LogRecord
}

func newInstrumentation(log *logger.Logger, record *logger.LogRecord) *instrumentation {
	return &instrumentation{
		logger: log,
		record: record.NewRecord().Scoped("ListDescendantsRuleSet"),
	}
}

func (i *instrumentation) setContext(ctx context.Context) {
	if traceId, ok := ctx.Value(contextKey.TraceId).(string); ok {
		i.record = i.record.TraceId(traceId)
	}
	if correlationId, ok := ctx.Value(contextKey.CorrelationId).(string); ok {
		i.record = i.record.CorrelationId(correlationId)
	}
}

func (i *instrumentation) setMetadata(metadata metadata) {
	i.record = i.record.Metadata(metadata)
}

func (i *instrumentation) startListingDescendantsRuleSet() {
	i.startedAt = time.Now()
	i.record = i.record.MessageObject("Starting listing descendants rule sets", "")
	i.doLog("startListingDescendantsRuleSet")
}

func (i *instrumentation) finishListingDescendantsRuleSet() {
	i.record.Duration(int(time.Since(i.startedAt))).MessageObject("finished listing descendants rule set", "")
	i.doLog("finishListingDescendantsRuleSet")
}

func (i *instrumentation) failedListingDescendantsRuleSet(err error) {
	i.record = i.record.MessageObject(
		"[VS] Error: fetching descendants rule sets from db",
		logger.Exception{
			ExceptionClass:   "ListDescendantsRuleSet Execute",
			Stacktrace:       "app/listDescendantsRuleSet/instrumentation.go failedListingDescendantsRuleSet",
			ExceptionMessage: err,
		},
	)

	i.doLog("failedListingDescendantsRuleSet")
}

func (i *instrumentation) doLog(loggerName string) {
	i.logger.Output.WithField("mdc", i.record.Mdc).WithField("message", i.record.Message).Info(loggerName)
}
