package dataobject

import (
	"fmt"

	"github.com/chanyut/promptpay/emvco/constants"
)

type MerchantName struct {
	name string
}

func (o MerchantName) ID() string {
	return "59"
}

func (o MerchantName) Length() int {
	return len(o.Value())
}

func (o MerchantName) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o MerchantName) Value() string {
	return o.name
}

func (o MerchantName) WithName(v string) MerchantName {
	if len(v) > 25 {
		panic(fmt.Errorf("name could be longer than 25"))
	}
	o.name = v
	return o
}
