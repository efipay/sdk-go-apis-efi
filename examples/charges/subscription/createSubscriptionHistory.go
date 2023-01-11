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
		"description": "This subscription was not fully paid",
	}

	res, err := efi.CreateSubscriptionHistory(1, body) // no lugar do 1 coloque o subscription_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}