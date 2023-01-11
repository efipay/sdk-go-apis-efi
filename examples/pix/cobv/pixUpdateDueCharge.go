package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)
	
	const txid = "adssshdsjdsjeyccdyddsasdstxid01"

	body := map[string]interface{} {
		"valor" : map[string]interface{}{
			"original" : "1.00",
		},
		"solicitacaoPagador" : "Enter the order number or identifier.",
	}

	res, err := efi.PixUpdateDueCharge(txid, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}