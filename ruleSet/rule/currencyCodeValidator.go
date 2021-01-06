package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type currencyCodeValidator struct {
	comparator compare.StringComparator
}

func newCurrencyCodeValidator(operator Operator, value string) (*currencyCodeValidator, error) {
	if !transaction.IsCurrencyCode(value) {
		return nil, InvalidValueError
	}

	var comparator compare.StringComparator

	switch operator {
	case OperatorEqual:
		comparator = compare.EqualString(value)
	case OperatorNotEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &currencyCodeValidator{comparator}, nil
}

func (v currencyCodeValidator) Validate(transaction transaction.Transaction) bool {
	return v.comparator(string(transaction.CurrencyCode))
}
