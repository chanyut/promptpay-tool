package dataobject

import "github.com/chanyut/promptpay/emvco/constants"

// TipOrConveTipOrConvenienceIndicator ff present, the Tip or Convenience Indicator shall contain a value of "01", "02" or "03".
// All other values are RFU.
// - A value of "01" shall be used if the mobile application should prompt the
// consumer to enter a tip to be paid to the merchant.
// - A value of "02" shall be used to indicate inclusion of the data object Value of
// Convenience Fee Fixed (ID "56").
// - A value of “03” shall be used to indicate inclusion of the data object Value of
// Convenience Fee Percentage (ID “57”).
// Note that even if the Transaction Amount is not present in the QR Code, this data
// object may still be present.
type TipOrConvenienceIndicator struct {
	value constants.TipOrConvenienceIndicatorValue
}

func (o TipOrConvenienceIndicator) ID() string {
	return "55"
}

func (o TipOrConvenienceIndicator) Length() int {
	return 2
}

func (o TipOrConvenienceIndicator) Format() constants.DataObjectFormat {
	return constants.N
}

func (o TipOrConvenienceIndicator) Value() string {
	return string(o.value)
}

func (o TipOrConvenienceIndicator) WithValue(v constants.TipOrConvenienceIndicatorValue) TipOrConvenienceIndicator {
	o.value = v
	return o
}
