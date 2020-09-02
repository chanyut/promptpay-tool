package promptpay

import (
	"fmt"

	"github.com/chanyut/promptpay/emvco"
	"github.com/chanyut/promptpay/emvco/dataobject"
)

type PromptpayMerchantPresentedQRCode struct {
	// promptpayID...
	promptpayIDType PromptpayIDType
	promptpayID     string

	amountString string
	currency     string
	countryCode  string

	dataObjects []dataobject.DataObject
}

func MakePromptpayMerchantPresentedQRCode() PromptpayMerchantPresentedQRCode {
	return PromptpayMerchantPresentedQRCode{}
}

func (p PromptpayMerchantPresentedQRCode) ID(idType PromptpayIDType, id string) PromptpayMerchantPresentedQRCode {
	if p.promptpayIDType == MobileNumber && len(id) != 13 {
		panic(fmt.Errorf("mobile number length must be 13"))
	} else if p.promptpayIDType == NationalID && len(id) != 13 {
		panic(fmt.Errorf("national ID length must be 13"))
	} else if p.promptpayIDType == TaxID && len(id) != 13 {
		panic(fmt.Errorf("tax ID length must be 13"))
	} else if p.promptpayIDType == EWalletID && len(id) != 15 {
		panic(fmt.Errorf("ewallet ID length must be 15"))
	} else if p.promptpayIDType == BankAccount && len(id) > 43 {
		panic(fmt.Errorf("bacnk account length could not be longer than 43"))
	}

	p.promptpayIDType = idType
	p.promptpayID = id
	return p
}

func (p PromptpayMerchantPresentedQRCode) MobileNumber(mobileNo string) PromptpayMerchantPresentedQRCode {
	return p.ID(MobileNumber, mobileNo)
}

func (p PromptpayMerchantPresentedQRCode) NationalID(natID string) PromptpayMerchantPresentedQRCode {
	return p.ID(NationalID, natID)
}

func (p PromptpayMerchantPresentedQRCode) TaxID(taxID string) PromptpayMerchantPresentedQRCode {
	return p.ID(TaxID, taxID)
}

func (p PromptpayMerchantPresentedQRCode) EWalletID(wID string) PromptpayMerchantPresentedQRCode {
	return p.ID(EWalletID, wID)
}

func (p PromptpayMerchantPresentedQRCode) BankAccount(acc string) PromptpayMerchantPresentedQRCode {
	return p.ID(BankAccount, acc)
}

// AmountF F => float
func (p PromptpayMerchantPresentedQRCode) AmountF(v float64) PromptpayMerchantPresentedQRCode {
	return p.AmountS(fmt.Sprintf("%.2f", v))
}

// AmountS S => string
func (p PromptpayMerchantPresentedQRCode) AmountS(v string) PromptpayMerchantPresentedQRCode {
	p.amountString = v
	return p
}

func (p PromptpayMerchantPresentedQRCode) Currency(v string) PromptpayMerchantPresentedQRCode {
	p.currency = v
	return p
}

func (p PromptpayMerchantPresentedQRCode) CountryCode(v string) PromptpayMerchantPresentedQRCode {
	p.countryCode = v
	return p
}

func (p *PromptpayMerchantPresentedQRCode) Build() {
	var promptpayID dataobject.DataObject
	if p.promptpayIDType == MobileNumber {
		promptpayID = MobileNumberDataObject{}.WithNumber(p.promptpayID)
	} else if p.promptpayIDType == NationalID || p.promptpayIDType == TaxID {
		promptpayID = NationalOrTaxIDDataObject{}.WithNationalOrTaxID(p.promptpayID)
	} else if p.promptpayIDType == EWalletID {
		promptpayID = EWalletDataObject{}.WithWalletID(p.promptpayID)
	} else if p.promptpayIDType == BankAccount {
		promptpayID = BankAccountDataObject{}.WithAccount(p.promptpayID)
	}

	// merchant account information
	merchantAccountInfo := dataobject.
		MerchantAccountInformation{}.
		WithID("29").WithValue([]dataobject.DataObject{
		dataobject.AIDDataObject{}.
			WithAppID(string(AppIDPromptpayMerchantPresented)),
		promptpayID,
	})

	p.dataObjects = []dataobject.DataObject{
		dataobject.PayloadFormatIndicator{},
		dataobject.PointOfInitiationMethod{},
		merchantAccountInfo,
		dataobject.TransactionAmount{}.WithAmountS(p.amountString),
		dataobject.CountryCode{}.WithCountryCode(p.countryCode),
		dataobject.TransactionCurrency{}.WithCurrencyCode(p.currency),
	}

	// calculate crc
	s := ""
	for _, o := range p.dataObjects {
		s += dataobject.Serialize(o)
	}
	crc := dataobject.CRC{}
	s = fmt.Sprintf("%s%s%02d", s, crc.ID(), crc.Length())
	crc = crc.WithValue(fmt.Sprintf("%X", emvco.CRC16XModem([]byte(s), 0xFFFF)))

	p.dataObjects = append(p.dataObjects, crc)
}

func (p PromptpayMerchantPresentedQRCode) ToString() string {
	s := ""
	for _, o := range p.dataObjects {
		s += dataobject.Serialize(o)
	}
	return s
}
