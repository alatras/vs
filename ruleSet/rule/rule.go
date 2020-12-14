package rule

import (
	"errors"

	"bitbucket.verifone.com/validation-service/transaction"
)

var (
	InvalidPropertyError = errors.New("invalid property while constructing validator")
	InvalidOperatorError = errors.New("invalid operator while constructing validator")
	InvalidValueError    = errors.New("invalid value while constructing validator")
)

type Property string

const (
	PropertyAmount                Property = "amount"
	PropertyCurrencyCode          Property = "currencyCode"
	PropertyCustomerCountryCode   Property = "customerCountryCode"
	PropertyCard                  Property = "card"
	PropertyIssuerCountryCode     Property = "issuerCountryCode"
	PropertyEntityId              Property = "entityId"
	PropertyCustomerId            Property = "customerId"
	PropertyCustomerIP            Property = "customerIP"
	PropertyCustomerIPCountryCode Property = "customerIPCountryCode"
)

type Operator string

const (
	OperatorLess           Operator = "<"
	OperatorLessOrEqual    Operator = "<="
	OperatorEqual          Operator = "=="
	OperatorNotEqual       Operator = "!="
	OperatorGreaterOrEqual Operator = ">="
	OperatorGreater        Operator = ">"
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
	var validator Validator
	var err error

	switch metadata.Property {
	case PropertyAmount:
		validator, err = newAmountValidator(metadata.Operator, metadata.Value)
	case PropertyCurrencyCode:
		validator, err = newCurrencyCodeValidator(metadata.Operator, metadata.Value)
	case PropertyCustomerCountryCode:
		validator, err = newCustomerCountryCodeValidator(metadata.Operator, metadata.Value)
	case PropertyCard:
		validator, err = newCardValidator(metadata.Operator, metadata.Value)
	case PropertyIssuerCountryCode:
		validator, err = newIssuerCountryCodeValidator(metadata.Operator, metadata.Value)
	case PropertyEntityId:
		validator, err = newEntityIdValidator(metadata.Operator, metadata.Value)
	case PropertyCustomerId:
		validator, err = newCustomerIdValidator(metadata.Operator, metadata.Value)
	case PropertyCustomerIP:
		validator, err = newCustomerIPValidator(metadata.Operator, metadata.Value)
	case PropertyCustomerIPCountryCode:
		validator, err = newCustomerIPCountryCodeValidator(metadata.Operator, metadata.Value)
	default:
		return nil, InvalidPropertyError
	}

	if err != nil {
		return nil, err
	}

	return validator, nil
}
