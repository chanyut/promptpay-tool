package constants

import "regexp"

type DataObjectFormat string

const (
	// N represents "Numeric"
	N DataObjectFormat = "N"
	// ANS represents "Alphanumeric special"
	ANS DataObjectFormat = "ANS"
	// S represents "String"
	S DataObjectFormat = "S"
)

func (f DataObjectFormat) Validate(v string) bool {
	if f == N {
		return regexp.MustCompile("[0-9]+").Match([]byte(v))
	} else if f == ANS {
		return regexp.MustCompile("[0-9]+").Match([]byte(v))
	} else if f == S {
		return true
	}
	return false
}

type TipOrConvenienceIndicatorValue string

const (
	PromptUserToEnter TipOrConvenienceIndicatorValue = "01"
	FixedValue        TipOrConvenienceIndicatorValue = "02"
	PercentageValue   TipOrConvenienceIndicatorValue = "03"
)
