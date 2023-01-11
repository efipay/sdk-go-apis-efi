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
		"email": "oldbuck@gerencianet.com.br",
	}

	res, err := efi.SendLinkEmail(1, body) //no lugar do 1 informe o id da cobrança desejada

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}