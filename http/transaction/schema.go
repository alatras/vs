package transaction

import (
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
)

type validateTransactionPayload struct {
	Amount       int    `json:"amount"`
	Organization string `json:"organization"`
}

type Metadata struct {
	Key      string   `json:"key"`
	Operator ruleSet.Operator `json:"operator"`
	Value    string   `json:"value"`
}

type RuleSet struct {
	Name     string `json:"name"`
	Metadata []Metadata `json:"rules"`
}


type validateTransactionResponse struct {
	Action report.Action `json:"action"`
	BlockedRuleSets []ruleSet.RuleSet `json:"block"`
	TaggedRuleSets  []ruleSet.RuleSet `json:"tags"`
}
