package tests

import (
	"validation-service/app/updateRuleSet"
	"validation-service/ruleSet"
	"validation-service/ruleSet/rule"
)

var mockUpdateRules = []updateRuleSet.Rule{
	{
		Key:      "amount",
		Operator: ">=",
		Value:    "1000",
	},
}

var mockInvalidOperatorRules = []updateRuleSet.Rule{
	{
		Key:      "amount",
		Operator: "!!!",
		Value:    "1000",
	},
}

var mockInvalidKeyRules = []updateRuleSet.Rule{
	{
		Key:      "invalid_field",
		Operator: ">=",
		Value:    "1000",
	},
}

var mockInvalidValueRules = []updateRuleSet.Rule{
	{
		Key:      "amount",
		Operator: "<=",
		Value:    "fff",
	},
}

var mockInvalidNoValueRules = []updateRuleSet.Rule{
	{
		Key:      "amount",
		Operator: "<=",
		Value:    "",
	},
}

var mockRuleSet = ruleSet.From(
	"1",
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
)

var mockUpdatedRuleSet = ruleSet.From(
	"1",
	"12345",
	"Test1",
	ruleSet.Block,
	[]rule.Metadata{
		{
			Property: "amount",
			Operator: ">=",
			Value:    "1000",
		},
	},
)
