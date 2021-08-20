package createRuleSet

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
		logger: log,
		record: record.NewRecord().Scoped("CreateRuleSet"),
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

func (i *instrumentation) startCreatingRuleSet() {
	i.startedAt = time.Now()
	i.record = i.record.MessageObject("Starting creating a rule set", "")
	i.doLog("startCreatingRuleSet")
}

func (i *instrumentation) invalidAction(action string) {
	i.record = i.record.MessageObject(
		"[VS] Error: Invalid action provided",
		logger.Exception{
			ExceptionClass:   "createRuleSet Execute",
			Stacktrace:       "app/createRuleSet/instrumentation.go invalidAction",
			ExceptionMessage: "Invalid action provided: " + action,
		},
	)
	i.doLog("invalidAction")
}

func (i *instrumentation) rulesetCreationFailed(error error) {
	i.record = i.record.MessageObject(
		"[VS] Error: ruleSet creation failed in repository",
		logger.Exception{
			ExceptionClass:   "createRuleSet Execute",
			Stacktrace:       "app/createRuleSet/instrumentation.go rulesetCreationFailed",
			ExceptionMessage: error,
		},
	)
	i.doLog("rulesetCreationFailed")
}

func (i *instrumentation) ruleMetadataInvalid(metadata rule.Metadata, error error) {
	i.record = i.record.Metadata(metadata).MessageObject(
		"[VS] Error: rule metadata is invalid",
		logger.Exception{
			ExceptionClass:   "createRuleSet Execute",
			Stacktrace:       "app/createRuleSet/instrumentation.go ruleMetadataInvalid",
			ExceptionMessage: error,
		},
	)
	i.doLog("rulesetCreationFailed")
}

func (i *instrumentation) finishCreatingRuleSet(ruleset ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)

	i.logger.Output.
		WithField("duration", duration).
		WithField("ruleSet", ruleset).
		Info("Finished creating a rule set")
}

func (i *instrumentation) doLog(loggerName string) {
	i.logger.Output.WithField("mdc", i.record.Mdc).WithField("message", i.record.Message).Info(loggerName)
}
