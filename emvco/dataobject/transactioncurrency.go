package dataobject

import (
	"fmt"

	"github.com/chanyut/promptpay/emvco/constants"
)

// TransactionCurrency indicates the currency code of the transaction.
// A 3-digit numeric value, as defined by [ISO 4217]. This value will be used
// by the mobile application to display a recognizable currency to the
// consumer whenever an amount is being displayed or whenever the
// consumer is prompted to enter an amount
type TransactionCurrency struct {
	currencyCode string
}

func (TransactionCurrency) ID() string {
	return "53"
}

func (TransactionCurrency) Length() int {
	return 3
}

func (TransactionCurrency) Format() constants.DataObjectFormat {
	return constants.N
}

func (o TransactionCurrency) Value() string {
	return o.currencyCode
}

func (o TransactionCurrency) WithCurrencyCode(v string) TransactionCurrency {
	if len(v) != 3 {
		panic(fmt.Errorf("invalid currency code: %v", v))
	}
	o.currencyCode = v
	return o
}
