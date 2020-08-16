package dataobject

import (
	"strconv"

	"github.com/chanyut/promptpay/emvco/constants"
	"github.com/pkg/errors"
)

type MerchantAccountInformation struct {
	id     string
	length int
	value  interface{}
}

func (o MerchantAccountInformation) ID() string {
	return o.id
}

func (o MerchantAccountInformation) Length() int {
	return o.length
}

func (MerchantAccountInformation) Format() constants.DataObjectFormat {
	return constants.ANS
}

func (o MerchantAccountInformation) Value() string {
	value := ""
	if s, ok := o.value.(string); ok {
		value = s
	} else if obj, ok := o.value.(DataObject); ok {
		value = obj.Value()
	} else {
		panic("invalid value type")
	}
	return value
}

func (o MerchantAccountInformation) WithID(id string) MerchantAccountInformation {
	v, err := strconv.Atoi(o.id)
	if err != nil {
		panic(errors.Wrapf(err, "invalid id: %v", id))
	}
	if v < 2 || v > 51 {
		panic(errors.New("id must be range [2, 51]"))
	}

	o.id = id
	return o
}

func (o MerchantAccountInformation) WithValue(v interface{}) MerchantAccountInformation {
	o.value = v
	if s, ok := v.(string); ok {
		o.length = len(s)
	} else if obj, ok := v.(DataObject); ok {
		o.length = len(obj.Value())
	} else {
		panic("invalid value type")
	}
	return o
}
