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
	const id = "a29"

	body := map[string]interface{} {		
		"valor": "0.01",
	}

	res, err := efi.PixDevolution(e2eid,id,body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}