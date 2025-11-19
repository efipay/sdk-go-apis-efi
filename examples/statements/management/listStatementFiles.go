package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/statements"
)

func main() {

	credentials := configs.Credentials
	efi := statements.NewEfiPay(credentials)

	res, err := efi.ListStatementFiles(nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
