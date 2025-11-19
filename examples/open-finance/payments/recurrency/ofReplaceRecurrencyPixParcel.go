package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	const identificadorPagamento = "urn:gerencianet:ea807997-9c28-47a7-8ebc-eeb672ea59f0"
	const endToEndId = "E0000000000000000000000000000000"

	body := map[string]interface{}{
		"valor": "9.99",
	}

	res, err := efi.OfReplaceRecurrencyPixParcel(identificadorPagamento, endToEndId, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
