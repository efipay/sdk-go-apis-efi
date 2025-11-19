package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/payments"
)

func main() {

	credentials := configs.Credentials
	efi := payments.NewEfiPay(credentials)

	body := map[string]interface{}{
		"url": "",
	}

	res, err := efi.PayDeleteWebhook(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
