package promptpay

type ApplicationID string

const (
	AppIDPromptpayMerchantPresented ApplicationID = "A000000677010111"
	AppIDPromptpayConsumerPresented ApplicationID = "A000000677010114"
)

type PromptpayIDType int

const (
	MobileNumber PromptpayIDType = iota
	NationalID
	TaxID
	EWalletID
	BankAccount
	Email
	BillerID
)

type MerchantCity string

const (
	Thailand MerchantCity = "TH"
)

type TransactionCurrency int

const (
	THB TransactionCurrency = 764
)
