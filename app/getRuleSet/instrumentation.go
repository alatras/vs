package getRuleSet

import (
	"context"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
	"validation-service/ruleSet"
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
		record: record.NewRecord().Scoped("GetRuleSet"),
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

func (i *instrumentation) startFetchingRuleSet() {
	i.startedAt = time.Now()
	i.record = i.record.MessageObject("Starting listing ancestors rule sets", "")
	i.doLog("startListingAncestorsRuleSet")
}

func (i *instrumentation) ruleSetFetchFailed(error error) {
	i.record = i.record.MessageObject(
		"[VS] Error: failed to fetch a rule set from the repository",
		logger.Exception{
			ExceptionClass:   "getRuleSet Execute",
			Stacktrace:       "app/getRuleSet/instrumentation.go ruleSetFetchFailed",
			ExceptionMessage: error,
		},
	)

	i.doLog("ruleSetFetchFailed")
}

func (i *instrumentation) ruleSetNotFound() {
	i.record = i.record.MessageObject("A rule set was not found", "")
	i.doLog("ruleSetNotFound")
}

func (i *instrumentation) finishFetchingRuleSet(ruleSet *ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)

	i.record.Duration(int(duration)).Data(
		map[string]interface{}{"ruleSet": *ruleSet},
	).MessageObject("Finished fetching a rule set", "")

	i.doLog("finishFetchingRuleSet")
}

func (i *instrumentation) doLog(loggerName string) {
	i.logger.Output.WithField("mdc", i.record.Mdc).WithField("message", i.record.Message).Info(loggerName)
}
