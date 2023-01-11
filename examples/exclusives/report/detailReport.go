package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const id = "cacf4fbf-09bc-44f5-bc32-143ffb2a37ff"

	res, err := efi.DetailReport(id) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}