package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/src/efipay/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)


	res, err := efi.GetAccountBalance(nil) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}