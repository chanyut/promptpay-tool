package dataobject

import "github.com/chanyut/promptpay/emvco/constants"

// PayloadFormatIndicator defines the version of the QR Code template and
// hence the conventions on the identifiers, lengths, and values.
// In this version of the specification, the Payload Format Indicator
// has the value "01".
type PayloadFormatIndicator struct{}

func (PayloadFormatIndicator) ID() string {
	return "00"
}

func (PayloadFormatIndicator) Length() int {
	return 2
}

func (PayloadFormatIndicator) Format() constants.DataObjectFormat {
	return constants.N
}

func (PayloadFormatIndicator) Value() string {
	return "01"
}
