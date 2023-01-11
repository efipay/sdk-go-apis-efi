package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/src/efipay/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)
	
	const e2eid = "E00416968202105211756Rh0iPsaJ1RK"

	res, err := efi.PixSendDetail(e2eid) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}