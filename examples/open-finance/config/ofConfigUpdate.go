package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	body := map[string]interface{} {
		"redirectURL" : "https://your-domain.com.br/redirect/",
		"webhookURL" : "https://your-domain.com.br/webhook/",
	}

	res, err := efi.OfConfigUpdate(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}