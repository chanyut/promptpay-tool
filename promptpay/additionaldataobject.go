package promptpay

import "github.com/chanyut/promptpay/emvco/constants"

type MobileNumberDataObject struct {
	number string
}

func (o MobileNumberDataObject) ID() string {
	return "01"
}

func (o MobileNumberDataObject) Length() int {
	return 13
}

func (o MobileNumberDataObject) Format() constants.DataObjectFormat {
	return constants.N
}

func (o MobileNumberDataObject) Value() string {
	return o.number
}

func (o MobileNumberDataObject) WithNumber(number string) MobileNumberDataObject {
	o.number = number
	return o
}

type NationalOrTaxIDDataObject struct {
	nationalOrTaxID string
}

func (o NationalOrTaxIDDataObject) ID() string {
	return "02"
}

func (o NationalOrTaxIDDataObject) Length() int {
	return 13
}

func (o NationalOrTaxIDDataObject) Format() constants.DataObjectFormat {
	return constants.N
}

func (o NationalOrTaxIDDataObject) Value() string {
	return o.nationalOrTaxID
}

func (o NationalOrTaxIDDataObject) WithNationalOrTaxID(id string) NationalOrTaxIDDataObject {
	o.nationalOrTaxID = id
	return o
}

type EWalletDataObject struct {
	walletID string
}

func (o EWalletDataObject) ID() string {
	return "03"
}

func (o EWalletDataObject) Length() int {
	return 15
}

func (o EWalletDataObject) Format() constants.DataObjectFormat {
	return constants.N
}

func (o EWalletDataObject) Value() string {
	return o.walletID
}

func (o EWalletDataObject) WithWalletID(id string) EWalletDataObject {
	o.walletID = id
	return o
}

type BankAccountDataObject struct {
	account string
}

func (o BankAccountDataObject) ID() string {
	return "04"
}

func (o BankAccountDataObject) Length() int {
	return len(o.Value())
}

func (o BankAccountDataObject) Format() constants.DataObjectFormat {
	return constants.N
}

func (o BankAccountDataObject) Value() string {
	return o.account
}

func (o BankAccountDataObject) WithAccount(account string) BankAccountDataObject {
	o.account = account
	return o
}
