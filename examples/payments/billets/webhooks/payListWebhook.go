package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/payments"
)

func main() {

	credentials := configs.Credentials
	efi := payments.NewEfiPay(credentials)

	params := map[string]string{
		"dataInicio": "2023-01-01T00:00:00Z",
		"dataFim":    "2024-12-31T23:59:59Z",
	}

	res, err := efi.PayListWebhook(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
