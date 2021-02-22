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
}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	return &instrumentation{
		logger: logger.Scoped("GetRuleSet"),
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

func (i *instrumentation) startFetchingRuleSet() {
	i.startedAt = time.Now()
	i.logger.Output.Info("Starting fetching a rule set")
}

func (i *instrumentation) ruleSetFetchFailed(error error) {
	i.logger.Output.
		WithError(error).
		Error("Failed to fetch a rule set from the repository")
}

func (i *instrumentation) ruleSetNotFound() {
	i.logger.Output.
		Error("A rule set was not found")
}

func (i *instrumentation) finishFetchingRuleSet(ruleSet *ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)

	i.logger.Output.
		WithField("duration", duration).
		WithField("ruleSet", *ruleSet).
		Info("Finished fetching a rule set")
}
