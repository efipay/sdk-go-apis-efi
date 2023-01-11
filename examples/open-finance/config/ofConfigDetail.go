package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	res, err := efi.OfConfigDetail() 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}