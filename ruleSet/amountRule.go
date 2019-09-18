package ruleSet

import (
	"bitbucket.verifone.com/validation-service/transaction"
)

type amountRule struct {
	validate func(int) bool
}

func newAmountRule(operator Operator, value int) (*amountRule, error) {
	f, err := newNumericValidator(operator, value)

	if err != nil {
		return nil, err
	}

	return &amountRule{
		validate: f,
	}, nil
}

func (rule *amountRule) Eval(transaction transaction.Transaction) bool {
	return rule.validate(transaction.Amount)
}
