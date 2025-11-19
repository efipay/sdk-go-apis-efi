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
		"id": "1",
	}

	body := map[string]interface{}{
		"cobsV": []interface{}{
			map[string]interface{}{
				"calendario": map[string]interface{}{
					"dataDeVencimento": "2025-11-28",
				},
				"txid": "fb2761260e550od593c7926bjb5cb650",
				"valor": map[string]interface{}{
					"original": "100.50",
				},
			},
			map[string]interface{}{
				"calendario": map[string]interface{}{
					"dataDeVencimento": "2020-12-31",
				},
				"txid": "7978c0c97ea84ppe78e8849634473c9f1",
				"valor": map[string]interface{}{
					"original": "110.50",
				},
			},
		},
	}

	res, err := efi.PixUpdateDueChargeBatch(params, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
