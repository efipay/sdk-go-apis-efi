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
		"identificador": "diario", // "diario", "semanal", "mensal"
	}

	body := map[string]interface{}{
		"periodicidade":     "diario",
		"status":            "ativo",
		"envia_email":       true,
		"comprimir_arquivo": true,
	}
	res, err := efi.UpdateStatementRecurrency(params, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
