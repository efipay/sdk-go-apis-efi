package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	
	const e2eid = "E00416968202105211756Rh0iPsaJ1RK"
	const id = ""

	res, err := efi.PixDetailDevolution(e2eid, id) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}