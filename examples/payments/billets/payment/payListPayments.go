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
		"dataInicio" : "2022-01-01",
		"dataFim" : "2024-12-31",
	}

	res, err := efi.PayListPayments(params) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}