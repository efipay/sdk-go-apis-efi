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
		"inicio": "2024-06-01T00:00:00Z",
		"fim":    "2024-12-31T23:59:59Z",
	}

	res, err := efi.MedList(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
