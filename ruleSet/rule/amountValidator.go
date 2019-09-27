package rule

import (
	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
	"strconv"
)

type amountValidator struct {
	amountComparator compare.Uint64Comparator
}

func newAmountValidator(operator operator, value string) (*amountValidator, error) {
	amount, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		return nil, errors.New("invalid value format")
	}

	var amountComparator compare.Uint64Comparator

	switch operator {
	case less:
		amountComparator = compare.LessThanUint64(amount)
	case lessOrEqual:
		amountComparator = compare.LessThanOrEqualUint64(amount)
	case equal:
		amountComparator = compare.EqualUint64(amount)
	case notEqual:
		amountComparator = compare.NotEqualUint64(amount)
	case greaterOrEqual:
		amountComparator = compare.GreaterThanOrEqualUint64(amount)
	case greater:
		amountComparator = compare.GreaterThanUint64(amount)
	default:
		return nil, errors.New("invalid operator")
	}

	return &amountValidator{amountComparator}, nil
}

func (v amountValidator) Validate(transaction transaction.Transaction) bool {
	return v.amountComparator(transaction.Amount)
}
