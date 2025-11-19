package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const id = ""

	res, err := efi.PixSplitDetailConfig(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
