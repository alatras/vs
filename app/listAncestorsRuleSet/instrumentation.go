package listAncestorsRuleSet

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
}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	return &instrumentation{
		logger: logger.Scoped("ListAncestorsRuleSet"),
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

func (i *instrumentation) startListingAncestorsRuleSet() {
	i.startedAt = time.Now()
	i.logger.Output.Info("starting listing ancestors rule sets")
}

func (i *instrumentation) finishListingAncestorsRuleSet() {
	i.logger.Output.
		WithField("duration", time.Since(i.startedAt)).
		Info("finished listing ancestors rule set")
}

func (i *instrumentation) failedListingAncestorsRuleSet(err error) {
	i.logger.Output.Logger.
		WithError(err).
		Error("error fetching ancestors rule sets from db")
}
