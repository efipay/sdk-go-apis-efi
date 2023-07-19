package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	const nome = ""

	res, err := efi.OfListParticipants(nome)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}