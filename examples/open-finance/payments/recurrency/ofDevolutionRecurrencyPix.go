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

	body := map[string]interface{}{
		"transacao1": map[string]interface{}{
			"endToEndId": "E09089356202408281500624f423208f",
			"valor":      "0.01",
		},
		"transacao2": map[string]interface{}{
			"endToEndId": "E09089356202408281500624f423208f",
			"valor":      "0.01",
		},
	}

	res, err := efi.OfDevolutionRecurrencyPix(identificadorPagamento, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
