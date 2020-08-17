package dataobject

import (
	"errors"

	"github.com/chanyut/promptpay/emvco/constants"
)

type PostalCode struct {
	code string
}

func (o PostalCode) ID() string {
	return "61"
}

func (o PostalCode) Length() int {
	return len(o.Value())
}

func (o PostalCode) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o PostalCode) Value() string {
	return o.code
}

func (o PostalCode) WithCode(v string) PostalCode {
	if len(v) > 10 {
		panic(errors.New("postal code could not longer than 10"))
	}
	o.code = v
	return o
}
