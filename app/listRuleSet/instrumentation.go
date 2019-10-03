package listRuleSet

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
		logger: logger.Scoped("ListRuleSet"),
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

func (i *instrumentation) startListingRuleSet() {
	i.startedAt = time.Now()
	i.logger.Output.Info("starting listing the rule sets")
}

func (i *instrumentation) finishListingRuleSet() {
	i.logger.Output.
		WithField("duration", time.Since(i.startedAt)).
		Info("finished listing rule set")
}

func (i *instrumentation) failedListingRuleSet(err error) {
	i.logger.Output.Logger.
		WithError(err).
		Error("error fetching ruleset from db")
}
