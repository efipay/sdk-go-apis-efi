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
		"name": "My plan",
		"interval": 2,
		"repeats": nil,
	}

	res, err := efi.CreatePlan(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}