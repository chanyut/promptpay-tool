package dataobject

import "github.com/chanyut/promptpay/emvco/constants"

type DataObject interface {
	ID() string
	Length() int
	Format() constants.DataObjectFormat
	Value() string
}
