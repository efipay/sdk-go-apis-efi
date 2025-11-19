package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/statements"
)

func main() {

	credentials := configs.Credentials
	efi := statements.NewEfiPay(credentials)

	body := map[string]interface{}{
		"periodicidade":     "diario",
		"envia_email":       true,
		"comprimir_arquivo": true,
	}

	res, err := efi.CreateStatementRecurrency(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
