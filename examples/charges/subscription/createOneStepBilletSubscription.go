package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	customer := map[string]interface{}{
		"name": "Gorbadoc Oldbuck",
		"cpf": "73504596295",
		"phone_number": "5144916523",
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
			"banking_billet": map[string]interface{} {
				"expire_at": "2022-12-12",
				"customer": customer,
			},
		},
	}

	res, err := efi.CreateOneStepSubscription(1, body) // no lugar do 1 coloque o plan_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}