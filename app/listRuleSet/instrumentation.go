package listRuleSet

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
		record: record.NewRecord().Scoped("ListRuleSet"),
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

func (i *instrumentation) startListingRuleSet() {
	i.startedAt = time.Now()
	i.record = i.record.MessageObject("Starting listing the rule sets", "")
	i.doLog("startListingRuleSet")
}

func (i *instrumentation) finishListingRuleSet() {
	i.record.Duration(int(time.Since(i.startedAt))).MessageObject("finished listing rule set", "")
	i.doLog("finishListingRuleSet")
}

func (i *instrumentation) failedListingRuleSet(err error) {
	i.record = i.record.MessageObject(
		"[VS] Error: fetching ruleset from db",
		logger.Exception{
			ExceptionClass:   "listRuleSet Execute",
			Stacktrace:       "app/listRuleSet/instrumentation.go failedListingRuleSet",
			ExceptionMessage: err,
		},
	)

	i.doLog("failedListingRuleSet")
}

func (i *instrumentation) doLog(loggerName string) {
	i.logger.Output.WithField("mdc", i.record.Mdc).WithField("message", i.record.Message).Info(loggerName)
}
