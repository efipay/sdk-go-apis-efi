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
		"x-skip-mtls-checking": "true",
	}

	body := map[string]interface{}{
		"webhookUrl": "https://usuario.recebedor.com/api/webhookrec/",
	}

	res, err := efi.PixConfigWebhookRecurrenceAutomatic(params, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
