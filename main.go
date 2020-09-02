package main

import (
	"log"

	"github.com/chanyut/promptpay/promptpay"
)

func main() {
	log.Println("started...")
	p := promptpay.PromptpayMerchantPresentedQRCode{}.
		AmountF(1520.00).
		MobileNumber("0066869743931").
		CountryCode("TH").
		Currency("764")
	p.Build()
	log.Println(p.ToString())
}
