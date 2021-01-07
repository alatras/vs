package report

import (
	"validation-service/ruleSet"
)

type Action string

const (
	Pass  Action = "PASS"
	Block Action = "BLOCK"
)

type Report struct {
	Action          Action            `json:"action"`
	BlockedRuleSets []ruleSet.RuleSet `json:"block"`
	TaggedRuleSets  []ruleSet.RuleSet `json:"tags"`
}

func New() Report {
	return Report{
		Action:          Pass,
		BlockedRuleSets: []ruleSet.RuleSet{},
		TaggedRuleSets:  []ruleSet.RuleSet{},
	}
}

func (report *Report) AppendBlockRuleSet(rs ruleSet.RuleSet) {
	report.Action = Block
	report.BlockedRuleSets = append(report.BlockedRuleSets, rs)
}

func (report *Report) AppendTagRuleSet(rs ruleSet.RuleSet) {
	report.TaggedRuleSets = append(report.TaggedRuleSets, rs)
}
