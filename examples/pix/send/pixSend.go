package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idEnvio = "0S000000000000000000000000000000000"

	body := map[string]interface{} {
		"valor": "0.01",
		"pagador": map[string]interface{} {
			"chave": "",
			"infoPagador": "Segue o pagamento da conta",
		},
		"favorecido": map[string]interface{} {			
			"chave": "",
		},
	}

	res, err := efi.PixSend(idEnvio, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}