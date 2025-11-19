package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	const inicio = "2022-03-01T03:01:35Z"
	const fim = "2022-12-05T22:01:35Z"	

	res, err := efi.OfListPixPayment(inicio, fim) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}