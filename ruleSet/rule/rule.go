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

type Property string

const (
	amount                Property = "amount"
	currencyCode          Property = "currencyCode"
	customerCountryCode   Property = "customerCountryCode"
	card                  Property = "card"
	issuerCountryCode     Property = "issuerCountryCode"
	entityId              Property = "entityId"
	customerId            Property = "customerId"
	customerIP            Property = "customerIP"
	customerIPCountryCode Property = "customerIPCountryCode"
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
	var validator Validator
	var err error

	/* properties to validate
	amount
	customer country code
	currency code

	-- To do --
	card - equal, not equal
	issuer country - equal, not equal
	entity - equal, not equal

	customer - equal, not equal
	customer IP - equal, not equal
	customer IP country - equal, not equal
	*/
	switch metadata.Property {
	case amount:
		validator, err = newAmountValidator(metadata.Operator, metadata.Value)
	case currencyCode:
		validator, err = newCurrencyCodeValidator(metadata.Operator, metadata.Value)
	case customerCountryCode:
		validator, err = newCustomerCountryCodeValidator(metadata.Operator, metadata.Value)
	case card:
		validator, err = newCardValidator(metadata.Operator, metadata.Value)
	case issuerCountryCode:
		validator, err = newIssuerCountryCodeValidator(metadata.Operator, metadata.Value)
	case entityId:
		validator, err = newEntityIdValidator(metadata.Operator, metadata.Value)
	case customerId:
		validator, err = newCustomerIdValidator(metadata.Operator, metadata.Value)
	case customerIP:
		validator, err = newCustomerIPValidator(metadata.Operator, metadata.Value)
	case customerIPCountryCode:
		validator, err = newCustomerIPCountryCodeValidator(metadata.Operator, metadata.Value)
	default:
		return nil, InvalidPropertyError
	}

	if err != nil {
		return nil, err
	}

	return validator, nil
}
