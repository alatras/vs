package validateTransaction

import (
	"context"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
	"validation-service/report"
)

type metadata = logger.Metadata

type instrumentation struct {
	logger    *logger.Logger
	createdAt time.Time
	startedAt time.Time
	record    *logger.LogRecord
}

func newInstrumentation(log *logger.Logger, record *logger.LogRecord) *instrumentation {
	return &instrumentation{
		logger:    log,
		createdAt: time.Now(),
		record:    record.NewRecord().Scoped("ValidateTransaction"),
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

func (i *instrumentation) startTransactionValidation() {
	delay := time.Since(i.createdAt)
	i.startedAt = time.Now()
	i.record = i.record.Delay(int(delay))
	i.record = i.record.MessageObject("Starting transaction validation", "")
	i.doLog("startTransactionValidation")
}

func (i *instrumentation) endTransactionValidation(r report.Report) {
	duration := time.Since(i.startedAt)
	i.record = i.record.Duration(int(duration))
	i.record = i.record.MessageObject("End transaction validation", r)
	i.doLog("endTransactionValidation")
}

func (i *instrumentation) doLog(loggerName string) {
	i.logger.Output.WithField("mdc", i.record.Mdc).WithField("message", i.record.Message).Info(loggerName)
}
