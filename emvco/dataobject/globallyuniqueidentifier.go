package dataobject

import "github.com/chanyut/promptpay/emvco/constants"

// GloballyUniqueIdentifier would be included in Merchant Account Information template
// it shall be used when the payment system corresponding to the Merchant Account Information is explicitly identified
// in the template.
//
// If present, a Merchant Account Information template shall contain a primitive
// Globally Unique Identifier data object with a data object ID "00", as defined in
// Table 4.2.
//
// The value of this data object shall contain one of the following:
//
// + An Application Identifier (AID) consisting of a RID registered with ISO and,
//   optionally, a PIX, as defined by [ISO 7816-4]. For example, "D840000000".
//
// + A [UUID] without the hyphen (-) separators. For example,
//   “581b314e257f41bfbbdc6384daa31d16”.
//
// + A reverse domain name. For example, “com.merchant.name”
type GloballyUniqueIdentifier struct {
	uniqueID string
}

func (o GloballyUniqueIdentifier) ID() string {
	return "00"
}

func (o GloballyUniqueIdentifier) Length() int {
	return len(o.Value())
}

func (o GloballyUniqueIdentifier) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o GloballyUniqueIdentifier) Value() string {
	return o.uniqueID
}

func (o GloballyUniqueIdentifier) WithUniqueID(id string) GloballyUniqueIdentifier {
	o.uniqueID = id
	return o
}
