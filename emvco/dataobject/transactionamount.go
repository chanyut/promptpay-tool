package dataobject

import (
	"fmt"
	"strconv"

	"github.com/chanyut/promptpay/emvco/constants"
	"github.com/pkg/errors"
)

// TransactionAmount Absent if the mobile application is to prompt the consumer
// to enter the transaction amount
// If present, the Transaction Amount shall be different from zero, shall only include
// (numeric) digits "0" to "9" and may contain a single "." character as the decimal
// mark. When the amount includes decimals, the "." character shall be used to
// separate the decimals from the integer value and the "." character may be present
// even if there are no decimals.
//
// The following are examples of valid Transaction Amounts: "98.73", "98" and "98.".
// The following are NOT valid Transaction Amounts: "98,73" and "3 705".
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
	if _, err := strconv.Atoi(v); err != nil {
		panic(errors.Wrapf(err, "invalid value: %s", v))
	}
	o.amountAsString = v
	return o
}
