package test

import (
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

var mockRuleSets = [1]ruleSet.RuleSet{ruleSet.New(
	"12345",
	"Test",
	ruleSet.Block,
	[]rule.Metadata{
		{
			Property: "amount",
			Operator: ">=",
			Value:    "1000",
		},
	},
	"TEST TAG",
)}
