package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	body := map[string]interface{}{
		"pagador": map[string]interface{}{
			"cpf":            "01234567890",
			"cnpj":           "01234567890123",
			"nome":           "Josué Real",
			"idParticipante": "64c189e5-9d4a-4319-aa5e-d02c36e1815d",
		},
		"favorecido": map[string]interface{}{
			"contaBanco": map[string]interface{}{
				"nome":        "Lucas Silva",
				"documento":   "12345678901",
				"codigoBanco": "09089356",
				"agencia":     "0001",
				"conta":       "654984",
				"tipoConta":   "TRAN",
			},
		},
		"assinatura": map[string]interface{}{
			"expiracao": "2026-08-27",
			"descricao": "Mensalidades do curso XYZ",
			"idProprio": "6236574863254",
			"configuracao": map[string]interface{}{
				"automatico": map[string]interface{}{
					"intervalo":          "ANUAL",
					"dataInicio":         "2025-06-06",
					"valorFixo":          "500.00",
					"valorMinimo":        "450.00",
					"valorMaximo":        "750.00",
					"permiteRetentativa": "",
					"primeiroPagamento": map[string]interface{}{
						"data":        "2024-06-08",
						"valor":       "9.99",
						"infoPagador": "Parcela 1",
					},
				},
			},
		},
	}

	res, err := efi.OfCreateAutomaticEnrollment(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
