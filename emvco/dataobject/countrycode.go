package dataobject

import "github.com/chanyut/promptpay/emvco/constants"

type CountryCode struct {
	code string
}

func (o CountryCode) ID() string {
	return "58"
}

func (o CountryCode) Length() int {
	return 2
}

func (o CountryCode) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o CountryCode) Value() string {
	return o.code
}

func (o CountryCode) WithCountryCode(v string) CountryCode {
	o.code = v
	return o
}
