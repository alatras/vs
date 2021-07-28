package test

import (
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

var mockRuleSet = ruleSet.New(
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
)
