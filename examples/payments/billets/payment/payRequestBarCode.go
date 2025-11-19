package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/payments"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	efi := payments.NewEfiPay(credentials)

	params := map[string]string {
		"codBarras" : "36400000000000000000000000000000000000000000000",
	}

	body := map[string]interface{} {
		"valor" : 500,
		"dataPagamento" : "2023-01-01",
		"descricao" : "Payment of the test billet",
	}

	res, err := efi.PayRequestBarCode(params, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}