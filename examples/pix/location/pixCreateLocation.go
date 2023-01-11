package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)



	body := map[string]interface{} {		
		"tipoCob": "cob",
	}

	res, err := efi.PixCreateLocation(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}