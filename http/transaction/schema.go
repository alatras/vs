package transaction

import (
	"bitbucket.verifone.com/validation-service/report"
	"bitbucket.verifone.com/validation-service/ruleSet"
)

type validateTransactionPayload struct {
	Amount       int    `json:"amount"`
	Organization string `json:"organization"`
}

type metadata struct {
	Key      string   `json:"key"`
	Operator ruleSet.Operator `json:"operator"`
	Value    string   `json:"value"`
}

type ruleSetResponse struct {
	Name     string `json:"name"`
	Metadata []metadata `json:"rules"`
}


type validateTransactionResponse struct {
	Action report.Action `json:"action"`
	BlockedRuleSets []ruleSetResponse `json:"block"`
	TaggedRuleSets  []ruleSetResponse `json:"tags"`
}
