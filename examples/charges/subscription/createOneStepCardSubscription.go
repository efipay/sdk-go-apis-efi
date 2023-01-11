package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	paymentToken := "428d7f3b2dc49117552ace464078557832c4ae4f";

	customer := map[string]interface{}{
		"name": "Gorbadoc Oldbuck",
		"cpf": "73504596295",
		"phone_number": "5144916523",
	}

	billingAddress := map[string]interface{} {
		"street": "Av JK",
		"number": 909,
		"neighborhood": "Bauxita",
		"zipcode": "35400000",
		"city": "Ouro Preto",
		"state": "MG",
	}

	body := map[string]interface{} {
		"items": []map[string]interface{}{
			{
				"name": "Product 1",
				"value": 1000,
				"amount": 2,
			},
		},
		"payment": map[string]interface{} {
			"credit_card": map[string]interface{} {
				"billing_address": billingAddress,
				"payment_token": paymentToken,
				"customer": customer,
			},
		},
	}

	res, err := efi.CreateOneStepSubscription(93454, body) // no lugar do 1 coloque o plan_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}