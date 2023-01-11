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
		"email": "gorbadoc.oldbuck@gerencianet.com.br",
	}

	res, err := efi.SendSubscriptionLinkEmail(1, body) // no lugar do 1 coloque o charge_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}