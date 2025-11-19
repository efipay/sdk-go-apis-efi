package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{}{
		"tipo": "PIX_RECEBIDO", //PIX_RECEBIDO, PIX_ENVIADO, DEVOLUCAO_RECEBIDA, DEVOLUCAO_ENVIADA
		"e2eids": []string{
			"E09089356202412261300API229e5352",
			"E09089356202412261300API3149af57",
		},
	}

	res, err := efi.PixResendWebhook(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
