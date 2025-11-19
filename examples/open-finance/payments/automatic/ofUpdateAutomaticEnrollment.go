package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	headers := map[string]string{
		"x-idempotency-key": "et sedaute sint officiapariatur amet tute sum",
	}

	body := map[string]interface{}{
		"identificador":  "urn:efi:19ba4105-9ae2-4637-89f2-96506d3c8770",
		"nomeFavorecido": "Marco Antonio de Brito",
		"status":         "revogado",
		"dataExpiracao":  "2021-05-21",
		"valorMaximo":    "4.22",
	}
	res, err := efi.OfUpdateAutomaticEnrollment(body, headers)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
