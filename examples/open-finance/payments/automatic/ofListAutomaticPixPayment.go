package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	params := map[string]interface{}{
		"identificadorAdesao": "urn:efi:19ba4105-9ae2-4637-89f2-96506d3c8770",
		"endToEndId":          "E00038166201907261559y6j6",
	}

	res, err := efi.OfListAutomaticPixPayment(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
