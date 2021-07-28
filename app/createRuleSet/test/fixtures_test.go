package test

import (
	"validation-service/app/createRuleSet"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

var mockRules = []createRuleSet.Rule{
	{
		Key:      "amount",
		Operator: ">=",
		Value:    "1000",
	},
}

var mockInvalidOperatorRules = []createRuleSet.Rule{
	{
		Key:      "amount",
		Operator: "!!!",
		Value:    "1000",
	},
}

var mockInvalidKeyRules = []createRuleSet.Rule{
	{
		Key:      "invalid_field",
		Operator: ">=",
		Value:    "1000",
	},
}

var mockInvalidValueRules = []createRuleSet.Rule{
	{
		Key:      "amount",
		Operator: "<=",
		Value:    "fff",
	},
}

var mockInvalidNoValueRules = []createRuleSet.Rule{
	{
		Key:      "amount",
		Operator: "<=",
		Value:    "",
	},
}

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
