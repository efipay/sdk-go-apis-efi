package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	
	const txid = "adssshdsjdsjeyccdyddsasdstxid2901"

	body := map[string]interface{} {
		
		"calendario": map[string]interface{} {
				"expiracao": 3600,
			},
		"devedor": map[string]interface{}{
			
				"cnpj": "12345678000195",
				"nome": "Empresa de Serviços SA",
			
		},
		"valor": map[string]interface{} {
			
				"original": "00.01",
			
		},
		"chave": "",
		"solicitacaoPagador": "Teste.",
		"infoAdicionais": []map[string]interface{} {
			{
				"nome": "Campo 1",
				"valor": "Informação Adicional1 do PSP-Recebedor",
			},
		},
	}

	res, err := efi.CreateCharge(txid,body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}