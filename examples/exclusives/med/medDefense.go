package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	params := map[string]string{
		"idInfracao": "00000000-0000-0000-0000-000000000000",
	}

	body := map[string]interface{}{
		"analise":       "aceito",
		"justificativa": "Justificativa",
	}
	res, err := efi.MedDefense(params, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
