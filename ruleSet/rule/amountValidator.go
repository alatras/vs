package rule

import (
	"strconv"

	"bitbucket.verifone.com/validation-service/ruleSet/compare"
	"bitbucket.verifone.com/validation-service/transaction"
)

type amountValidator struct {
	validator func(trxAmount uint64, minorUnits int) bool
}

func newAmountValidator(operator Operator, value string) (*amountValidator, error) {
	var amountComparator func(uint64) compare.Uint64Comparator

	switch operator {
	case OperatorLess:
		amountComparator = compare.LessThanUint64
	case OperatorLessOrEqual:
		amountComparator = compare.LessThanOrEqualUint64
	case OperatorEqual:
		amountComparator = compare.EqualUint64
	case OperatorNotEqual:
		amountComparator = compare.NotEqualUint64
	case OperatorGreaterOrEqual:
		amountComparator = compare.GreaterThanOrEqualUint64
	case OperatorGreater:
		amountComparator = compare.GreaterThanUint64
	default:
		return nil, InvalidOperatorError
	}

	compareAmount, err := strconv.ParseUint(value, 10, 64)

	if err != nil {
		return nil, InvalidValueError
	}

	validator := func(trxAmount uint64, minorUnits int) bool {
		compareAmountWithMinorUnits := compareAmount

		for i := 0; i < minorUnits; i++ {
			compareAmountWithMinorUnits *= 10
		}

		return amountComparator(compareAmountWithMinorUnits)(trxAmount)
	}

	return &amountValidator{validator}, nil
}

func (v amountValidator) Validate(transaction transaction.Transaction) bool {
	return v.validator(transaction.Amount, transaction.MinorUnits)
}
