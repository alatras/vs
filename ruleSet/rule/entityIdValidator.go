package rule

import (
	"validation-service/ruleSet/compare"
	"validation-service/transaction"
)

type entityIdValidator struct {
	comparator compare.StringComparator
}

func newEntityIdValidator(operator Operator, value string) (*entityIdValidator, error) {
	var comparator compare.StringComparator

	switch operator {
	case OperatorEqual:
		comparator = compare.EqualString(value)
	case OperatorNotEqual:
		comparator = compare.NotEqualString(value)
	default:
		return nil, InvalidOperatorError
	}

	return &entityIdValidator{comparator}, nil
}

func (e entityIdValidator) Validate(transaction transaction.Transaction) bool {
	return e.comparator(transaction.EntityId)
}
