package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{} {}	
	const key = "c40fe422-cfab-4dd4-b1f8-368db6f98c71" //chave pix a ser deletada

	res, err := efi.PixDeleteEvp(key, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}