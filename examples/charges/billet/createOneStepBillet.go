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
		"cpf": "60740262386",
		"phone_number": "51944916523",
		"email": "gorb.oldbuck@gerencianet.com.br",
	}

	body := map[string]interface{} {
		"payment": map[string]interface{} {
			"banking_billet": map[string]interface{} {
				"expire_at": "2023-12-12",
				"customer": customer,
			},
		},
		"items": []map[string]interface{}{
			{
				"name": "Product 1",
				"value": 500,
				"amount": 2,
			},
		},
	}

	res, err := efi.CreateOneStepCharge(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}