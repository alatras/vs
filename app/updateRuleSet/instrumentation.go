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
}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	return &instrumentation{
		logger: logger.Scoped("UpdateRuleSet"),
	}
}

func (i *instrumentation) setContext(ctx context.Context) {
	if traceId, ok := ctx.Value(contextKey.TraceId).(string); ok {
		i.logger = i.logger.WithTraceId(traceId)
	}
}

func (i *instrumentation) setMetadata(metadata metadata) {
	i.logger = i.logger.WithMetadata(metadata)
}

func (i *instrumentation) startUpdatingRuleSet() {
	i.startedAt = time.Now()
	i.logger.Output.Info("Started updating a rule set")
}

func (i *instrumentation) invalidAction(action string) {
	i.logger.Output.
		WithField("action", action).
		Error("Invalid action provided")
}

func (i *instrumentation) ruleSetUpdateFailed(error error) {
	i.logger.Output.
		WithError(error).
		Error("RuleSet update failed in repository")
}

func (i *instrumentation) ruleMetadataInvalid(metadata rule.Metadata, error error) {
	i.logger.Output.
		WithError(error).
		WithField("ruleMetadata", metadata).
		Error("Rule metadata is invalid")
}

func (i *instrumentation) finishUpdatingRuleSet(ruleset ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)

	i.logger.Output.
		WithField("duration", duration).
		WithField("ruleSet", ruleset).
		Info("Finished updating a rule set")
}

func (i *instrumentation) ruleSetNotReplaced(ruleset ruleSet.RuleSet) {
	i.logger.Output.
		WithField("ruleSet", ruleset).
		Info("RuleSet not updated")
}
