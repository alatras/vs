package deleteRuleSet

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
	startedAt time.Time
}

var fileLog = logger.FileLogger{}

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
	now := time.Now()
	entry := i.logger.Output
	i.startedAt = now
	msg := "Starting deleting a rule set"

	fileLog.StartedAt = now
	fileLog.Error = nil
	fileLog.Message = msg
	fileLog.Fields = []field{}

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleSetDeletionFailed(error error) {
	entry := i.logger.Output.WithError(error)
	msg := "Failed to delete a rule set in the repository"

	fileLog.Error = error
	fileLog.Message = msg
	fileLog.Fields = []field{}

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleSetNotFound() {
	entry := i.logger.Output
	msg := "A rule set was not found"

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleSetDeleted() {
	duration := time.Since(i.startedAt)
	entry := i.logger.Output.WithField("duration", duration)
	msg := "Finished deleting a rule set"

	fileLog.Message = msg
	fileLog.Fields = []field{{"duration", duration}}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}
