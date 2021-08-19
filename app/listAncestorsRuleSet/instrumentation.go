package listAncestorsRuleSet

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
		record: record.NewRecord().Scoped("listAncestorsRuleSet"),
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
	// i.logger = i.logger.WithMetadata(metadata)
}

func (i *instrumentation) startListingAncestorsRuleSet() {
	i.startedAt = time.Now()
	i.record = i.record.MessageObject("Starting listing ancestors rule sets", "")
	i.doLog(i.record.Mdc, i.record.Message, "startListingAncestorsRuleSet")
}

func (i *instrumentation) finishListingAncestorsRuleSet() {
	i.record.Duration(int(time.Since(i.startedAt))).MessageObject("Finished listing ancestors rule set", "")
	i.doLog(i.record.Mdc, i.record.Message, "finishListingAncestorsRuleSet")
}

func (i *instrumentation) failedListingAncestorsRuleSet(err error) {
	i.record = i.record.MessageObject(
		"[VS] Error: fetching ancestors rule sets from db",
		logger.Exception{
			ExceptionClass:   "listAncestorsRuleSet Execute",
			Stacktrace:       "app/listAncestorsRuleSet/instrumentation.go failedListingAncestorsRuleSet",
			ExceptionMessage: err,
		},
	)

	i.doLog(i.record.Mdc, i.record.Message, "failedListingAncestorsRuleSet")
}

func (i *instrumentation) doLog(
	mdc logger.MDC,
	message logger.Message,
	loggerName string,
) {
	i.logger.Output.WithField(
		"mdc", i.record.Mdc,
	).WithField(
		"message", i.record.Message,
	).Info(loggerName)
}
