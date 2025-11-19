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
			"idParticipante": "ebbed125-5cd7-42e3-965d-2e7af8e3b7ae",
			"cpf":            "90686175832",
			"cnpj":           "90293071000112",
		},
		"favorecido": map[string]interface{}{
			"contaBanco": map[string]interface{}{
				"nome":        "Lucas Silva",
				"documento":   "17558266300",
				"codigoBanco": "09089356",
				"agencia":     "0001",
				"conta":       "654984",
				"tipoConta":   "CACC",
			},
		},
		"pagamento": map[string]interface{}{
			"valor":            "9.90",
			"codigoCidadeIBGE": "3540000",
			"infoPagador":      "Churrasco",
			"idProprio":        "6236574863254",
			"dataAgendamento":  "2024-08-06",
		},
	}

	res, err := efi.OfStartSchedulePixPayment(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
