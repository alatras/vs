package rule

import (
	"strconv"
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type fraudScoreValidator struct {
	validator func(trxScore string) bool
}

func newFraudScoreValidator(operator Operator, ruleScore string) (*fraudScoreValidator, error) {
	ruleScoreParse, err := strconv.ParseFloat(ruleScore, 64)
	if err != nil || !(ruleScoreParse >= 0 && ruleScoreParse <= 1) {
		return nil, InvalidValueError
	}

	var amountComparator func(float64) compare.Float64Comparator

	switch operator {
	case OperatorLess:
		amountComparator = compare.LessThanFloat64
	case OperatorLessOrEqual:
		amountComparator = compare.LessThanOrEqualFloat64
	case OperatorEqual:
		amountComparator = compare.EqualFloat64
	case OperatorNotEqual:
		amountComparator = compare.NotEqualFloat64
	case OperatorGreaterOrEqual:
		amountComparator = compare.GreaterThanOrEqualFloat64
	case OperatorGreater:
		amountComparator = compare.GreaterThanFloat64
	default:
		return nil, InvalidOperatorError
	}

	validator := func(trxScore string) bool {
		trxScoreParse, err := strconv.ParseFloat(trxScore, 64)
		if err != nil || !(trxScoreParse >= 0 && trxScoreParse <= 1) {
			return false
		}

		return amountComparator(ruleScoreParse)(trxScoreParse)
	}

	return &fraudScoreValidator{validator}, nil
}

func (v fraudScoreValidator) Validate(transaction transaction.Transaction) bool {
	return v.validator(transaction.FraudScore)
}
