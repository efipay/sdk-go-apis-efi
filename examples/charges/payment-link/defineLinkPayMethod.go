package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	body := map[string]interface{} {
		"billet_discount": 1,
		"card_discount": 1,
		"message": "teste",
		"expire_at": "2022-12-30",
		"request_delivery_address": false,
		"payment_method": "all",
	}

	res, err := efi.DefineLinkPayMethod(1, body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}