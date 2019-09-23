package transaction

import (
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
)

type ValidateTransactionPayload struct {
	Amount int    `json:"amount"`
	Entity string `json:"entity"`
}

type Metadata struct {
	Key      string           `json:"key"`
	Operator ruleSet.Operator `json:"operator"`
	Value    string           `json:"value"`
}

type RuleSetResponse struct {
	Name     string     `json:"name"`
	Metadata []Metadata `json:"rules"`
}

type ValidateTransactionResponse struct {
	Action          report.Action     `json:"action"`
	BlockedRuleSets []RuleSetResponse `json:"block"`
	TaggedRuleSets  []RuleSetResponse `json:"tags"`
}
