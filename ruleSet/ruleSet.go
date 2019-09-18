package ruleSet

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"strconv"
)

type Action int

const (
	Pass Action = iota
	Block
	Tag
)

type Operator string

const (
	Less           Operator = "<"
	LessOrEqual    Operator = "<="
	Equal          Operator = "=="
	NotEqual       Operator = "!="
	GreaterOrEqual Operator = ">="
	Greater        Operator = ">"
)

type Metadata struct {
	Key      string   `json:"key"`
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
}

type rule interface {
	Eval(transaction transaction.Transaction) bool
}

type RuleSet struct {
	Name     string `json:"name"`
	action   Action
	Metadata []Metadata `json:"rules"`
	rules    []rule
}

func New(name string, action Action, metadata []Metadata) (RuleSet, error) {
	ruleSet := RuleSet{
		Name:     name,
		action:   action,
		Metadata: metadata,
		rules:    []rule{},
	}

	for _, m := range metadata {
		switch m.Key {
		case "amount":
			value, err := strconv.Atoi(m.Value)

			if err != nil {
				return ruleSet, err
			}

			rule, err := newAmountRule(m.Operator, value)

			if err != nil {
				return ruleSet, err
			}

			ruleSet.rules = append(ruleSet.rules, rule)
		}
	}

	return ruleSet, nil
}

func (r RuleSet) EvaluateTransaction(trx transaction.Transaction) Action {
	for _, rule := range r.rules {
		match := rule.Eval(trx)

		if !match {
			return Pass
		}
	}

	return r.action
}

type Repository interface {
	ListForOrganization(organization string) []RuleSet
}
