package validateTransaction

import (
	"context"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
)

type metadata = logger.Metadata
type field = logger.Field

type instrumentation struct {
	logger    *logger.Logger
	createdAt time.Time
	startedAt time.Time
}

var fileLog = logger.FileLogger{}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	scope := "ValidateTransaction"
	now := time.Now()

	fileLog.Scope = scope
	fileLog.CreatedAt = now

	return &instrumentation{
		logger:    logger.Scoped(scope),
		createdAt: now,
	}
}

func (i *instrumentation) setContext(ctx context.Context) {
	if traceId, ok := ctx.Value(contextKey.TraceId).(string); ok {
		i.logger = i.logger.WithTraceId(traceId)
		fileLog.TraceID = traceId
	}
}

func (i *instrumentation) setMetadata(metadata metadata) {
	i.logger = i.logger.WithMetadata(metadata)
	fileLog.Metadata = metadata
}

func (i *instrumentation) startTransactionValidation() {
	delay := time.Since(i.createdAt)
	i.startedAt = time.Now()

	entry := i.logger.Output.WithField("delay", delay)
	msg := "Starting transaction validation"

	fileLog.Message = msg
	fileLog.Fields = []field{{"delay", delay}}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) endTransactionValidation() {
	duration := time.Since(i.startedAt)
	msg := "Transaction validation finished"
	entry := i.logger.Output.WithField("duration", duration)

	fileLog.Message = msg
	fileLog.Fields = []field{{"duration", duration}}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}
