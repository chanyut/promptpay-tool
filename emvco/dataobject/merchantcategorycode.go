package dataobject

import (
	"errors"

	"github.com/chanyut/promptpay/emvco/constants"
)

// MerchantCategoryCode as defined by [ISO 18245] and assigned by the Acquirer
type MerchantCategoryCode struct {
	categoryCode string
}

func (MerchantCategoryCode) ID() string {
	return "52"
}

func (MerchantCategoryCode) Length() int {
	return 4
}

func (MerchantCategoryCode) Format() constants.DataObjectFormat {
	return constants.N
}

func (o MerchantCategoryCode) Value() string {
	return o.categoryCode
}

func (o MerchantCategoryCode) WithCategoryCode(v string) MerchantCategoryCode {
	if len(v) != 4 {
		panic(errors.New("invalid category code"))
	}
	o.categoryCode = v
	return o
}
