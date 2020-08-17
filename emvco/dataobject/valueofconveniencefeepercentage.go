package dataobject

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chanyut/promptpay/emvco/constants"
	"github.com/pkg/errors"
)

// ValueOfConvenienceFeePercentage ...
// The Value of Convenience Fee Percentage shall be present if the data object Tip
// or Convenience Indicator (ID "55") is present with a value of "03"and only values
// between “00.01” and “99.99” shall be used. Otherwise this data object shall be
// absent.
//
// + If present, the Value of Convenience Fee Percentage shall only include (numeric)
// digits "0" to "9" and may contain a single "." character as the decimal mark.
//
// + When the Value of the Convenience Fee Percentage includes decimals, the "."
// character shall be used to separate the decimals from the integer value and the "."
// character may be present even if there are no decimals.
//
// The Value of Convenience Fee Percentage shall not contain any other characters.
// For example, the “%” character must not be included.
// The above describes the only acceptable format for the Value of Convenience
// Fee Percentage.

type ValueOfConvenienceFeePercentage struct {
	amountAsString string
}

func (o ValueOfConvenienceFeePercentage) ID() string {
	return "57"
}

func (o ValueOfConvenienceFeePercentage) Length() int {
	return len(o.Value())
}

func (o ValueOfConvenienceFeePercentage) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o ValueOfConvenienceFeePercentage) Value() string {
	return o.amountAsString
}

func (o ValueOfConvenienceFeePercentage) WithAmountI(v int) ValueOfConvenienceFeePercentage {
	s := fmt.Sprintf("%d", v)
	return o.WithAmountS(s)
}

func (o ValueOfConvenienceFeePercentage) WithAmountS(v string) ValueOfConvenienceFeePercentage {
	const maxLength = 5
	if len(v) > maxLength {
		panic(fmt.Errorf("amount's length should not greater than %d", maxLength))
	}
	if _, err := strconv.Atoi(v); err != nil {
		panic(errors.Wrapf(err, "invalid value: %s", v))
	}
	if strings.Contains(v, "%") {
		panic(errors.New("value could not contain % symbol"))
	}
	o.amountAsString = v
	return o
}
