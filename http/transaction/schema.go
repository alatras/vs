package transaction

import (
	"validation-service/report"
)

type ValidateTransactionPayload struct {
	Transaction transaction `json:"transaction"`
}

type transaction struct {
	Amount     amount       `json:"amount"`
	Merchant   merchant     `json:"merchant"`
	Customer   customer     `json:"customer"`
	Instrument []instrument `json:"instrument"`
}

type amount struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currencyCode"`
}

type ValidateTransactionResponse struct {
	report.Report
}

type merchant struct {
	Organisation organisation `json:"organisation"`
}

type organisation struct {
	UUID string `json:"UUID"`
}

type customer struct {
	Country                string                 `json:"country"`
	CustomerIdentification customerIdentification `json:"identification"`
	IP                     string                 `json:"IPAddressV4"`
	IPCountry              string                 `json:"IPCountry"`
}

type customerIdentification struct {
	CustomerId string `json:"customerId"`
}

type instrument struct {
	Type       string `json:"instrumentType"`
	CardNumber string `json:"cardNumber"`
	Country    string `json:"country"`
}
