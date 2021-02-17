package updateRuleSet

import (
	"context"
	"errors"
	"time"
	"validation-service/enums/contextKey"
	"validation-service/logger"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

type metadata = logger.Metadata
type field = logger.Field

type instrumentation struct {
	logger    *logger.Logger
	startedAt time.Time
}

var fileLog = logger.FileLogger{}

func newInstrumentation(logger *logger.Logger) *instrumentation {
	scope := "UpdateRuleSet"
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

func (i *instrumentation) startUpdatingRuleSet() {
	now := time.Now()
	msg := "Started updating a rule set"
	entry := i.logger.Output
	i.startedAt = now

	fileLog.StartedAt = now
	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) invalidAction(action string) {
	entry := i.logger.Output.WithField("action", action)
	err := "invalid action provided"

	fileLog.Fields = []field{{"action", action}}
	fileLog.Message = ""
	fileLog.Error = errors.New(err)

	entry.Error(err)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleSetUpdateFailed(error error) {
	entry := i.logger.Output.WithError(error)
	msg := "RuleSet update failed in repository"

	fileLog.Error = error
	fileLog.Message = msg
	fileLog.Fields = []field{}

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleMetadataInvalid(metadata rule.Metadata, error error) {
	entry := i.logger.Output.WithError(error).WithField("ruleMetadata", metadata)
	msg := "Rule metadata is invalid"

	fileLog.Fields = []field{{"ruleMetadata", metadata}}
	fileLog.Message = msg
	fileLog.Error = error

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) finishUpdatingRuleSet(ruleset ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)
	entry := i.logger.Output.WithField("duration", duration).WithField("ruleSet", ruleset)
	msg := "Finished updating a rule set"

	fileLog.Fields = []field{{"duration", duration}, {"ruleSet", ruleset}}
	fileLog.Message = msg
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleSetNotReplaced(ruleset ruleSet.RuleSet) {
	entry := i.logger.Output.WithField("ruleSet", ruleset)
	msg := "RuleSet not updated"

	fileLog.Fields = []field{{"ruleSet", ruleset}}
	fileLog.Message = msg
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}
