package ruleSet

import (
	"errors"
)

func newNumericValidator(operator Operator, value int) (func(int) bool, error) {
	switch operator {
	case Less:
		return lessThan(value), nil
	case LessOrEqual:
		return lessThanOrEqual(value), nil
	case Equal:
		return equal(value), nil
	case NotEqual:
		return notEqual(value), nil
	case GreaterOrEqual:
		return greaterThanOrEqual(value), nil
	case Greater:
		return greaterThan(value), nil
	default:
		return nil, errors.New("invalid operator")
	}
}

func lessThan(value int) func(int) bool {
	return func(compare int) bool {
		return compare < value
	}
}

func lessThanOrEqual(value int) func(int) bool {
	return func(compare int) bool {
		return compare <= value
	}
}

func equal(value int) func(int) bool {
	return func(compare int) bool {
		return compare == value
	}
}

func notEqual(value int) func(int) bool {
	return func(compare int) bool {
		return compare != value
	}
}

func greaterThanOrEqual(value int) func(int) bool {
	return func(compare int) bool {
		return compare >= value
	}
}

func greaterThan(value int) func(int) bool {
	return func(compare int) bool {
		return compare > value
	}
}
