package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
)

var (
	InvalidPropertyError = errors.New("invalid property while constructing validator")
	InvalidOperatorError = errors.New("invalid operator while constructing validator")
	InvalidValueError    = errors.New("invalid value while constructing validator")
)

type property string

const (
	amount              property = "amount"
	currencyCode        property = "currencyCode"
	customerCountryCode property = "customerCountryCode"
)

type operator string

const (
	less           operator = "<"
	lessOrEqual    operator = "<="
	equal          operator = "=="
	notEqual       operator = "!="
	greaterOrEqual operator = ">="
	greater        operator = ">"
)

type Metadata struct {
	Property property `json:"key"`
	Operator operator `json:"operator"`
	Value    string   `json:"value"`
}

type Validator interface {
	Validate(trx transaction.Transaction) bool
}

func NewValidator(metadata Metadata) (Validator, error) {
	var validator Validator
	var err error

	switch metadata.Property {
	case amount:
		validator, err = newAmountValidator(metadata.Operator, metadata.Value)
	case currencyCode:
		validator, err = newCurrencyCodeValidator(metadata.Operator, metadata.Value)
	case customerCountryCode:
		validator, err = newCustomerCountryCodeValidator(metadata.Operator, metadata.Value)
	default:
		return nil, InvalidPropertyError
	}

	if err != nil {
		return nil, err
	}

	return validator, nil
}
