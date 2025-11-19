package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	params := map[string]string{
		"e2eid": "E0000000000000000000000000000",
		// "txid": "0000000000000000000000000000000",
		// "idEnvio": "0000000000000000000000000000000",
		// "rtrId": "D0000000000000000000000000000",
	}

	res, err := efi.PixGetReceipt(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
