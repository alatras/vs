package report

import (
	"bitbucket.verifone.com/validation-service/ruleSet"
)

type action string

const (
	Pass  action = "PASS"
	Block        = "BLOCK"
)

type Report struct {
	Action          action            `json:"action"`
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
