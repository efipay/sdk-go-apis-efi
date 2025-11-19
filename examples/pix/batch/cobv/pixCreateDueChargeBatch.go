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
		"id": "1",
	}

	body := map[string]interface{}{
		"descricao": "Cobranças dos alunos do turno vespertino",
		"cobsV": []interface{}{
			map[string]interface{}{
				"calendario": map[string]interface{}{
					"dataDeVencimento":       "2024-11-28",
					"validadeAposVencimento": 30,
				},
				"txid": "fb2761260e550od593c7926bjb5cb650",
				"devedor": map[string]interface{}{
					"cpf":  "01234567890",
					"nome": "João Souza",
				},
				"valor": map[string]interface{}{
					"original": "100.50",
				},
				"chave":              "a572c113-7d13-49ba-9988-0yy7e061a356",
				"solicitacaoPagador": "Informar matrícula",
			},
			map[string]interface{}{
				"calendario": map[string]interface{}{
					"dataDeVencimento":       "2020-12-31",
					"validadeAposVencimento": 30,
				},
				"txid": "7978c0c97ea84ppe78e8849634473c9f1",
				"devedor": map[string]interface{}{
					"cpf":  "15311295449",
					"nome": "Manoel Silva",
				},
				"valor": map[string]interface{}{
					"original": "100.50",
				},
				"chave":              "a572c113-7d13-49ba-9988-0yy7e061a356",
				"solicitacaoPagador": "Informar matrícula",
			},
		},
	}

	res, err := efi.PixCreateDueChargeBatch(params, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
