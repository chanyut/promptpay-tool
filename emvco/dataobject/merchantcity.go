package dataobject

import (
	"errors"

	"github.com/chanyut/promptpay/emvco/constants"
)

type MerchantCity struct {
	city string
}

func (o MerchantCity) ID() string {
	return "60"
}

func (o MerchantCity) Length() int {
	return len(o.Value())
}

func (o MerchantCity) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o MerchantCity) Value() string {
	return o.city
}

func (o MerchantCity) WithCity(city string) MerchantCity {
	if len(city) > 15 {
		panic(errors.New("city could not be longer than 15"))
	}
	o.city = city
	return o
}
