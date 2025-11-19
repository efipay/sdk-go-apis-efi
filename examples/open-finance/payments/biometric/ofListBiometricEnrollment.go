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
		"cpf": "90686175832",
		// 	// "cnpj": "",
	}

	res, err := efi.OfListBiometricEnrollment(requestParams)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
