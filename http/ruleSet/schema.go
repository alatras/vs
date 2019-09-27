package ruleSet

type RulePayload struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type GetRuleSetResponse struct {
	Id     string        `json:"id"`
	Entity string        `json:"entity"`
	Name   string        `json:"name"`
	Action string        `json:"action"`
	Rules  []RulePayload `json:"rules"`
}
