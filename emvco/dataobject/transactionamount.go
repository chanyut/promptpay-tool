package dataobject

import (
	"fmt"

	"github.com/chanyut/promptpay/emvco/constants"
)

type TransactionAmount struct {
	amountAsString string
}

func (TransactionAmount) ID() string {
	return "54"
}

func (o TransactionAmount) Length() int {
	return len(o.amountAsString)
}

func (o TransactionAmount) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o TransactionAmount) Value() string {
	return o.amountAsString
}

func (o TransactionAmount) WithAmountI(v int) TransactionAmount {
	s := fmt.Sprintf("%d", v)
	return o.WithAmountS(s)
}

func (o TransactionAmount) WithAmountS(v string) TransactionAmount {
	const maxLength = 13
	if len(v) > maxLength {
		panic(fmt.Errorf("amount's length should not greater than %d", maxLength))
	}
	o.amountAsString = v
	return o
}
