package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{} {
		"dataMovimento": "2022-04-24",
    	"tipoRegistros": map[string]interface{} {
			"pixRecebido": true,
			"pixEnviadoChave": true,
			"pixEnviadoDadosBancarios": true,
			"estornoPixEnviado": true,
			"pixDevolucaoEnviada": true,
			"pixDevolucaoRecebida": true,
			"tarifaPixEnviado": true,
			"tarifaPixRecebido": true,
			"estornoTarifaPixEnviado": true,
			"saldoDiaAnterior": true,
			"saldoDia": true,
			"transferenciaEnviada": true,
			"transferenciaRecebida": true,
			"estornoTransferenciaEnviada": true,
			"tarifaTransferenciaEnviada": true,
			"estornoTarifaTransferenciaEnviada": true,
    	},
	}

	
	res, err := efi.CreateReport(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}