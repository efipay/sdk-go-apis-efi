package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)
	
	headers := map[string]string{
		"x-idempotency-key": "et sedaute sint officiapariatur amet tute sum",
	} 
	

	body := map[string]interface{} {
		"pagador": map[string]interface{} {
				"idParticipante": "ebbed125-5cd7-42e3-965d-2e7af8e3b7ae",
				"cpf" : "90686175832",
		},
		"favorecido": map[string]interface{}{
				"contaBanco": map[string]interface{}{
					"codigoBanco" : "364",
					"agencia" : "0001",
					"documento" : "89696876378",
					"nome" : "Gorbadoc Oldbuck",
					"conta" : "000000",
					"tipoConta" : "CACC",
				},
		},
		"valor": "0.01",
		"infoPagador": "Teste.",
		"idProprio": "Order_00001",
	}

	res, err := efi.OfStartPixPayment(body, headers) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}