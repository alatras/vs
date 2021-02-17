package getRuleSet

import (
	"context"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
	"validation-service/ruleSet"
)

type metadata = logger.Metadata
type field = logger.Field

type instrumentation struct {
	logger    *logger.Logger
	startedAt time.Time
}

var fileLog = logger.FileLogger{}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	scope := "GetRuleSet"
	fileLog.Scope = scope

	return &instrumentation{
		logger: logger.Scoped(scope),
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

func (i *instrumentation) startFetchingRuleSet() {
	entry := i.logger.Output
	msg := "Starting fetching a rule set"
	now := time.Now()
	i.startedAt = now

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.StartedAt = now
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleSetFetchFailed(error error) {
	entry := i.logger.Output.WithError(error)
	msg := "Failed to fetch a rule set from the repository"

	fileLog.Fields = []field{}
	fileLog.Error = error
	fileLog.Message = msg

	entry.Error(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleSetNotFound() {
	i.logger.Output.
		Error("A rule set was not found")
}

func (i *instrumentation) finishFetchingRuleSet(ruleSet *ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)
	entry := i.logger.Output.WithField("duration", duration).WithField("ruleSet", *ruleSet)
	msg := "Finished fetching a rule set"

	fileLog.Fields = []field{{"duration", duration}, {"ruleSet", *ruleSet}}
	fileLog.Message = msg
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}
