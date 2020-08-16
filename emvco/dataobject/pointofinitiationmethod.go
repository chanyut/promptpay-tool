package dataobject

import "github.com/chanyut/promptpay/emvco/constants"

// PointOfInitiationMethod identifies the communication technology (here QR Code) and whether the
// data is static or dynamic.
// The Point of Initiation Method has a value of "11" for static QR Codes and
// a value of "12" for dynamic QR Codes.
// The value of "11" is used when the same QR Code is shown for more than
// one transaction.
// The value of "12" is used when a new QR Code is shown for each
// transaction
type PointOfInitiationMethod struct {
	isDynamic bool
}

func (PointOfInitiationMethod) ID() string {
	return "01"
}
func (PointOfInitiationMethod) Length() int {
	return 2
}
func (PointOfInitiationMethod) Format() constants.DataObjectFormat {
	return constants.N
}

func (o PointOfInitiationMethod) Value() string {
	if o.isDynamic {
		return "12"
	}
	return "11"
}

func (o PointOfInitiationMethod) SetDynamic(v bool) PointOfInitiationMethod {
	o.isDynamic = v
	return o
}
