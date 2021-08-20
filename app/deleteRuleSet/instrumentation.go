package deleteRuleSet

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
		record: record.NewRecord().Scoped("DeleteRuleSet"),
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

func (i *instrumentation) startDeletingRuleSet() {
	i.startedAt = time.Now()
	i.record = i.record.MessageObject("Starting deleting a rule set", "")
	i.doLog("startDeletingRuleSet")
}

func (i *instrumentation) ruleSetDeletionFailed(error error) {
	i.record = i.record.MessageObject(
		"[VS] Error: failed to delete a rule set in the repository",
		logger.Exception{
			ExceptionClass:   "deleteRuleSet Execute",
			Stacktrace:       "app/deleteRuleSet/instrumentation.go ruleSetDeletionFailed",
			ExceptionMessage: error,
		},
	)
	i.doLog("ruleSetDeletionFailed")
}

func (i *instrumentation) ruleSetNotFound() {
	i.record = i.record.MessageObject("A rule set was not found", "")
	i.doLog("ruleSetNotFound")
}

func (i *instrumentation) ruleSetDeleted() {
	duration := time.Since(i.startedAt)
	i.record.Duration(int(duration)).MessageObject("Finished deleting a rule set", "")
	i.doLog("ruleSetDeleted")
}

func (i *instrumentation) doLog(loggerName string) {
	i.logger.Output.WithField("mdc", i.record.Mdc).WithField("message", i.record.Message).Info(loggerName)
}
