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
		"pagador": map[string]interface{}{
			"cpf":            "90686175832",
			"idParticipante": "ebbed125-5cd7-42e3-965d-2e7af8e3b7ae",
		},
	}

	res, err := efi.OfCreateBiometricEnrollment(body, headers)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
