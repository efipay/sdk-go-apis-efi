package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	const identificadorPagamento = "urn:efi:ae71713f-875b-4af3-9d85-0bcb43288847"

	res, err := efi.OfCancelSchedulePix(identificadorPagamento)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
