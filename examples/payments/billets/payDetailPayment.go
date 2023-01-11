package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/payments"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	efi := payments.NewEfiPay(credentials)

	params := map[string]string {
		"idPagamento" : "1",
		// "status" : "REALIZADO", // EM_PROCESSAMENTO, AGENDADO, LIQUIDADO, CANCELADO, NAO_REALIZADO 
	}

	res, err := efi.PayDetailPayment(params) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}