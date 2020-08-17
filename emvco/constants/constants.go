package constants

type DataObjectFormat string

const (
	// N represents "Numeric"
	N DataObjectFormat = "N"
	// ANS represents "Alphanumeric special"
	ANS DataObjectFormat = "ANS"
	// S represents "String"
	S DataObjectFormat = "S"
)

type TipOrConvenienceIndicatorValue string

const (
	PromptUserToEnter TipOrConvenienceIndicatorValue = "01"
	FixedValue        TipOrConvenienceIndicatorValue = "02"
	PercentageValue   TipOrConvenienceIndicatorValue = "03"
)
