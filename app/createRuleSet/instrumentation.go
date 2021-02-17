package createRuleSet

import (
	"context"
	"time"
	"errors"
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
	return &instrumentation{
		logger: logger.Scoped("CreateRuleSet"),
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

func (i *instrumentation) startCreatingRuleSet() {
	entry := i.logger.Output
	now := time.Now()
	i.startedAt = now
	msg := "Starting creating a rule set"

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.StartedAt = now
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) invalidAction(action string) {
	entry := i.logger.Output.WithField("action", action)
	err := "invalid action provided"

	fileLog.Message = ""
	fileLog.Fields = []field{{"action", action}}
	fileLog.Error = errors.New(err)

	entry.Error(err)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) rulesetCreationFailed(error error) {
	entry := i.logger.Output.WithError(error)
	msg := "RuleSet creation failed in repository"

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.Error = error

	entry.Error("RuleSet creation failed in repository")
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) ruleMetadataInvalid(metadata rule.Metadata, error error) {
	entry := i.logger.Output.WithError(error).WithField("ruleMetadata", metadata)
	msg := "Rule metadata is invalid"

	fileLog.Message = msg
	fileLog.Fields = []field{}
	fileLog.Error = error

	entry.Error(msg)
	logger.WriteToFile(fileLog)
}

func (i *instrumentation) finishCreatingRuleSet(ruleset ruleSet.RuleSet) {
	duration := time.Since(i.startedAt)
	entry := i.logger.Output.WithField("duration", duration).WithField("ruleSet", ruleset)
	msg := "Finished creating a rule set"

	fileLog.Message = msg
	fileLog.Fields = []field{{"duration", duration}, {"ruleSet", ruleset}}
	fileLog.Error = nil

	entry.Info(msg)
	logger.WriteToFile(fileLog)
}
