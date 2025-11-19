package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	
	params := map[string]string {
		"inicio" : "2022-01-01T00:00:00Z",
		"fim" : "2024-12-31T23:59:59Z",
		// "txid" : "00000000000000000000000000000000000",
		// "txIdPresente" : true,
		// "devolucaoPresente" : false,
		// "cpf" : "12345678909", // Filtrar pelo CPF do pagador. Nao pode ser usado ao mesmo tempo que CNPJ.
		// "cnpj" : "12345678909", // Filtrar pelo CNPJ do pagador. Nao pode ser usado ao mesmo tempo que CPF.
		// "paginacao.paginaAtual" : 1,
		// "paginacao.itensPorPagina" : 10,
	}


	res, err := efi.PixReceivedList(params) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}