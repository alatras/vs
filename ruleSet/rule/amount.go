package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
	"strconv"
)

type amountValidator struct {
	amountComparator compare.IntComparator
}

func newAmountValidator(operator operator, value string) (*amountValidator, error) {
	amount, err := strconv.Atoi(value)

	if err != nil {
		return nil, err
	}

	var amountComparator compare.IntComparator

	switch operator {
	case less:
		amountComparator = compare.LessThanInt(amount)
	case lessOrEqual:
		amountComparator = compare.LessThanOrEqualInt(amount)
	case equal:
		amountComparator = compare.EqualInt(amount)
	case notEqual:
		amountComparator = compare.NotEqualInt(amount)
	case greaterOrEqual:
		amountComparator = compare.GreaterThanOrEqualInt(amount)
	case greater:
		amountComparator = compare.GreaterThanInt(amount)
	default:
		return nil, errors.New("invalid operator")
	}

	return &amountValidator{amountComparator}, nil
}

func (v amountValidator) Validate(transaction transaction.Transaction) bool {
	return v.amountComparator(transaction.Amount)
}
