package rule

import (
	"bitbucket.verifone.com/validation-service/transaction"
	"errors"
)

type Property string

const (
	amount Property = "amount"
)

type Operator string

const (
	less           Operator = "<"
	lessOrEqual    Operator = "<="
	equal          Operator = "=="
	notEqual       Operator = "!="
	greaterOrEqual Operator = ">="
	greater        Operator = ">"
)

type Metadata struct {
	Property Property `json:"key"`
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
}

type Validator interface {
	Validate(trx transaction.Transaction) bool
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
