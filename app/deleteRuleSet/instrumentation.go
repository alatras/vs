package deleteRuleSet

import (
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"bitbucket.verifone.com/validation-service/logger"
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
		logger: logger.Scoped("DeleteRuleSet"),
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

func (i *instrumentation) startDeletingRuleSet() {
	i.startedAt = time.Now()
	i.logger.Output.Info("Starting deleting a rule set")
}

func (i *instrumentation) ruleSetDeletionFailed(error error) {
	i.logger.Output.
		WithError(error).
		Error("Failed to delete a rule set in the repository")
}

func (i *instrumentation) ruleSetNotFound() {
	i.logger.Output.
		Error("A rule set was not found")
}

func (i *instrumentation) ruleSetDeleted() {
	duration := time.Since(i.startedAt)

	i.logger.Output.
		WithField("duration", duration).
		Info("Finished deleting a rule set")
}
