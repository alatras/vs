package createRuleSet

import (
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"bitbucket.verifone.com/validation-service/logger"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
	"context"
	"time"
)

type metadata = logger.Metadata

type instrumentation struct {
	logger    *logger.Logger
	startedAt time.Time
}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	return &instrumentation{
		logger: logger.Scoped("CreateRuleSet"),
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

func (i *instrumentation) startCreatingRuleset() {
	i.startedAt = time.Now()
	i.logger.Output.Info("Starting creating a rule set")
}

func (i *instrumentation) invalidAction(action string) {
	i.logger.Output.
		WithField("action", action).
		Error("Invalid action provided")
}

func (i *instrumentation) rulesetCreationFailed(error error) {
	i.logger.Output.
		WithError(error).
		Error("Ruleset creation failed in repository")
}

func (i *instrumentation) ruleMetadataInvalid(metadata rule.Metadata, error error) {
	i.logger.Output.
		WithError(error).
		WithField("ruleMetadata", metadata).
		Error("Rule metadata is invalid")
}

func (i *instrumentation) finishCreatingRuleset(ruleset ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)

	i.logger.Output.
		WithField("duration", duration).
		WithField("ruleSet", ruleset).
		Info("Finished creating a rule set")
}
