package listDescendantsRuleSet

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
		logger: logger.Scoped("ListDescendantsRuleSet"),
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

func (i *instrumentation) startListingDescendantsRuleSet() {
	i.startedAt = time.Now()
	i.logger.Output.Info("starting listing ancestors rule sets")
}

func (i *instrumentation) finishListingDescendantsRuleSet() {
	i.logger.Output.
		WithField("duration", time.Since(i.startedAt)).
		Info("finished listing ancestors rule set")
}

func (i *instrumentation) failedListingDescendantsRuleSet(err error) {
	i.logger.Output.Logger.
		WithError(err).
		Error("error fetching ancestors rule sets from db")
}

func (i *instrumentation) failedGetDescendants(err error) {
	i.logger.Output.Logger.
		WithError(err).
		Error("error fetching ancestors from the entity service")
}
