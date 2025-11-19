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
		"idEnvio": "",
	}

	body := map[string]interface{}{
		"valor": "0.01",
		"pagador": map[string]interface{}{
			"chave":       "",
			"infoPagador": "Segue o pagamento da conta",
		},
		"favorecido": map[string]interface{}{
			"chave": "",
		},
	}

	res, err := efi.PixSendSameOwnership(params, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
