package updateRuleSet

import (
	"context"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

type metadata = logger.Metadata

type instrumentation struct {
	logger    *logger.Logger
	startedAt time.Time
	record    *logger.LogRecord
}

func newInstrumentation(log *logger.Logger, record *logger.LogRecord) *instrumentation {
	return &instrumentation{
		logger:    log,
		startedAt: time.Now(),
		record:    record.NewRecord().Scoped("UpdateRuleset"),
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
	// i.logger = i.logger.WithMetadata(metadata)
	i.record = i.record.Metadata(metadata)
}

func (i *instrumentation) startUpdatingRuleSet() {
	i.startedAt = time.Now()
	i.record = i.record.MessageObject("Starting transaction validation", "")
	i.doLog(i.record.Mdc, i.record.Message, "startTransactionValidation")
}

func (i *instrumentation) invalidAction(action string) {
	i.record = i.record.MessageObject(
		"[VS] Error: Invalid action provided",
		logger.Exception{
			ExceptionClass:   "ruleSetUpdate Execute",
			Stacktrace:       "app/updateRuleset/instrumentation.go invalidAction",
			ExceptionMessage: "Invalid action provided: " + action,
		},
	)

	i.doLog(i.record.Mdc, i.record.Message, "invalidAction")
}

func (i *instrumentation) ruleSetUpdateFailed(error error) {
	i.record = i.record.MessageObject(
		"[VS] Error: RuleSet update failed in repository",
		logger.Exception{
			ExceptionClass:   "ruleSetUpdate Execute",
			Stacktrace:       "app/updateRuleset/instrumentation.go ruleSetUpdateFailed",
			ExceptionMessage: error,
		},
	)

	i.doLog(i.record.Mdc, i.record.Message, "invalidAction")
}

func (i *instrumentation) ruleMetadataInvalid(metadata rule.Metadata, error error) {
	i.record = i.record.Metadata(metadata).MessageObject(
		"[VS] Error: Rule metadata is invalid",
		logger.Exception{
			ExceptionClass:   "ruleSetUpdate Execute",
			Stacktrace:       "app/updateRuleset/instrumentation.go ruleMetadataInvalid",
			ExceptionMessage: error,
		},
	)

	i.doLog(i.record.Mdc, i.record.Message, "invalidAction")
}

func (i *instrumentation) finishUpdatingRuleSet(ruleset ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)
	i.record = i.record.Duration(int(duration)).Data(
		map[string]interface{}{"ruleSet": ruleset},
	).MessageObject("Finished updating a rule set", "")

	i.doLog(i.record.Mdc, i.record.Message, "finishUpdatingRuleSet")
	// i.logger.Output.
	// 	WithField("duration", duration).
	// 	WithField("ruleSet", ruleset).
	// 	Info("Finished updating a rule set")
}

func (i *instrumentation) ruleSetNotReplaced(ruleset ruleSet.RuleSet) {
	i.record.Data(
		map[string]interface{}{"ruleSet": ruleset},
	).MessageObject("RuleSet not updated", "")

	i.doLog(i.record.Mdc, i.record.Message, "ruleSetNotReplaced")
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
