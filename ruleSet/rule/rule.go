package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
)

type property string

const (
	amount property = "amount"
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
	IsMatch(trx transaction.Transaction) bool
}

func NewValidator(metadata Metadata) (Validator, error) {
	switch metadata.Property {
	case amount:
		validator, err := newAmountValidator(metadata.Operator, metadata.Value)
		if err != nil {
			return nil, err
		}
		return validator, nil
	default:
		return nil, errors.New("invalid property")
	}
}
