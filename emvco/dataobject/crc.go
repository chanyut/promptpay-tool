package dataobject

import (
	"github.com/chanyut/promptpay/emvco/constants"
	"github.com/pkg/errors"
)

type CRC struct {
	value string
}

func (o CRC) ID() string {
	return "63"
}

func (o CRC) Length() int {
	return 4
}

func (o CRC) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o CRC) Value() string {
	return o.value
}

func (o CRC) WithValue(v string) CRC {
	if len(v) != 4 {
		panic(errors.New("CRC value's length must be 4"))
	}
	o.value = v
	return o
}
