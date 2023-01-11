package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	body := map[string]interface{} {
		"title": "Balancete Demonstrativo",
		"body": []map[string]interface{} {
			{
				"header": "Demonstrativo de Consumo",
				"tables": []map[string]interface{} {
					{
						"rows": [][]map[string]interface{} {
							{
								{
									"align": "left",
									"color": "#000000",
									"style": "bold",
									"text": "Exemplo de despesa",
									"colspan": 2,
								},
								{
									"align": "left",
									"color": "#000000",
									"style": "bold",
									"text": "Total lançado",
									"colspan": 2,
								},
							},
							{
								{
									"align": "left",
									"color": "#000000",
									"style": "normal",
									"text": "Instalação",
									"colspan": 2,
								},
								{
									"align": "left",
									"color": "#000000",
									"style": "normal",
									"text": "R$ 100,00",
									"colspan": 2,
								},
							},
						},
					},
				},
			},
			{
				"header": "Balancete Geral",
				"tables": []map[string]interface{} {
					{
						"rows": [][]map[string]interface{} {
							{
								{
									"align": "left",
									"color": "#000000",
									"style": "normal",
									"text": "Confira na documentação da Gerencianet todas as configurações possíveis de um boleto balancete.",
									"colspan": 4,
								},
							},
						},
					},
				},
			},
		},	
	}

	res, err := efi.DefineBalanceSheetBillet(1, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}