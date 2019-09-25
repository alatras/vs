package test

import (
	"bitbucket.verifone.com/validation-service/app/createRuleSet"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"bitbucket.verifone.com/validation-service/ruleSet/rule"
)

var mockRules = []createRuleSet.Rule{
	{
		Key:      "amount",
		Operator: ">=",
		Value:    "1000",
	},
}

var (
	mockRuleSet = ruleSet.New("12345", "Test", ruleSet.Block, []rule.Metadata{
		{
			Property: "amount",
			Operator: ">=",
			Value:    "1000",
		},
	})
)
