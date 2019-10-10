package validateTransaction

import (
	"bitbucket.verifone.com/validation-service/enums/contextKey"
	"bitbucket.verifone.com/validation-service/logger"
	"context"
	"time"
)

type metadata = logger.Metadata

type instrumentation struct {
	logger    *logger.Logger
	createdAt time.Time
	startedAt time.Time
}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	return &instrumentation{
		logger:    logger.Scoped("ValidateTransaction"),
		createdAt: time.Now(),
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

func (i *instrumentation) startTransactionValidation() {
	delay := time.Since(i.createdAt)
	i.startedAt = time.Now()
	i.logger.Output.WithField("delay", delay).Info("Starting transaction validation")
}

func (i *instrumentation) endTransactionValidation() {
	duration := time.Since(i.startedAt)
	i.logger.Output.WithField("duration", duration).Info("Transaction validation finished")
}
