package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	requestParams := map[string]interface{}{
		"inicio": "2024-06-01T00:00:00Z",
		"fim":    "2024-09-14T23:59:59Z",
		"status": "",
		//"identificadorAdesao": "",
		//"idProprio": "",
		//"documento": "",
	}

	res, err := efi.OfListAutomaticEnrollment(requestParams)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
