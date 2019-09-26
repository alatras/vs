package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
)

type property string

const (
	amount      property = "amount"
	countryCode property = "countryCode"
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

type Operator interface {
	Operator() operator
}

func (o operator) Operator() operator {
	return o
}

type Metadata struct {
	Property property `json:"key"`
	Operator Operator `json:"operator"`
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
	case countryCode:
		validator, err = newCountryCodeValidator(metadata.Operator, metadata.Value)
	default:
		return nil, errors.New("invalid validator property")
	}

	if err != nil {
		return nil, err
	}

	return validator, nil
}
