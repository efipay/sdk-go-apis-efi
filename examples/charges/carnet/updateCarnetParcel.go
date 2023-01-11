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
		"expire_at": "2021-12-12",
	}

	res, err := efi.UpdateCarnetParcel(1, 1, body) // no lugar dos 1s coloque o carnet_id e o numero da parcela respectivamente

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}