package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	headers := map[string]string{
		"x-idempotency-key": "et sedaute sint officiapariatur amet tute sum",
	}

	body := map[string]interface{}{
		"identificadorVinculo": "enrollmentId",
		"favorecido": map[string]interface{}{
			"contaBanco": map[string]interface{}{
				"nome":        "GORBADOCK SILVA",
				"documento":   "01234567890",
				"codigoBanco": "09089356",
				"agencia":     "0001",
				"conta":       "654984",
				"tipoConta":   "TRAN",
			},
		},
		"pagamento": map[string]interface{}{
			"valor": "250.00",
		},
	}

	res, err := efi.OfCreateBiometricPixPayment(body, headers)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
