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
		"descricao": "Waask",
		"lancamento": map[string]interface{}{
			"imediato": true,
		},
		"split": map[string]interface{}{
			"divisaoTarifa": "assumir_total",
			"minhaParte": map[string]interface{}{
				"tipo":  "porcentagem",
				"valor": "85.00",
			},
			"repasses": []map[string]interface{}{
				{
					"tipo":  "porcentagem",
					"valor": "15.00",
					"favorecido": map[string]interface{}{
						"cpf":   "",
						"conta": "",
					},
				},
				{
					"tipo":  "porcentagem",
					"valor": "15.00",
					"favorecido": map[string]interface{}{
						"cpf":   "",
						"conta": "",
					},
				},
			},
		},
	}

	res, err := efi.PixSplitConfig(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
