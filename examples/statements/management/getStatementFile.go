package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/statements"
)

func main() {

	credentials := configs.Credentials
	efi := statements.NewEfiPay(credentials)

	params := map[string]string{
		"nome_arquivo": "nome_arquivo",
	}

	res, err := efi.GetStatementFile(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
