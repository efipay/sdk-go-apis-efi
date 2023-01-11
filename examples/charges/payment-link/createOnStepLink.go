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
		"items": []map[string]interface{}{
			{
				"name" : "Product One",
				"value": 600,
				"amount": 1,
			},
		},
		"settings": map[string]interface{}{
				"payment_method": "all",
				"expire_at": "2023-12-15",
				"request_delivery_address": false,
		},
	}

	res, err := efi.CreateOneStepLink(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}