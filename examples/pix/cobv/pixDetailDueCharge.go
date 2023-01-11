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

	

	res, err := efi.DetailDueCharge(txid) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}