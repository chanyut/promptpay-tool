package dataobject

import "github.com/chanyut/promptpay/emvco/constants"

type AIDDataObject struct {
	appID string
}

func (o AIDDataObject) ID() string {
	return "00"
}

func (o AIDDataObject) Length() int {
	return 16
}

func (o AIDDataObject) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o AIDDataObject) Value() string {
	return o.appID
}

func (o AIDDataObject) WithAppID(appID string) AIDDataObject {
	o.appID = appID
	return o
}
