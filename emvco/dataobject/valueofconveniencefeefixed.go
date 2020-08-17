package dataobject

import (
	"fmt"
	"strconv"

	"github.com/chanyut/promptpay/emvco/constants"
	"github.com/pkg/errors"
)

// ValueOfConvenienceFeeFixed ...
// The Value of Convenience Fee Fixed shall be present and different from zero
// if the data object Tip or Convenience Indicator (ID "55") is present with a value of "02".
// Otherwise this data object shall be absent.
//
// + If present, the Value of Convenience Fee Fixed shall only include (numeric) digits
// "0" to "9" and may contain a single "." character as the decimal mark.
//
// + When the Value of the Convenience Fee Fixed includes decimals, the "."
// character shall be used to separate the decimals from the integer value.
// The "." character may be present even if there are no decimals.
// The number of digits after the decimal mark should align with the currency
// exponent associated to the currency code defined in [ISO 4217].
type ValueOfConvenienceFeeFixed struct {
	amountAsString string
}

func (o ValueOfConvenienceFeeFixed) ID() string {
	return "56"
}

func (o ValueOfConvenienceFeeFixed) Length() int {
	return len(o.Value())
}

func (o ValueOfConvenienceFeeFixed) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o ValueOfConvenienceFeeFixed) Value() string {
	return o.amountAsString
}

func (o ValueOfConvenienceFeeFixed) WithAmountI(v int) ValueOfConvenienceFeeFixed {
	s := fmt.Sprintf("%d", v)
	return o.WithAmountS(s)
}

func (o ValueOfConvenienceFeeFixed) WithAmountS(v string) ValueOfConvenienceFeeFixed {
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
