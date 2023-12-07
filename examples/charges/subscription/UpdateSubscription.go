package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	customer := map[string]interface{}{
		"email": "gorbadoc.oldbuck@email.com",
		"phone_number": "5144916523",
	}

	items := []map[string]interface{}{
		{
			"name": "Product 1",
			"value": 1000,
			"amount": 2,
		},
	}

	shippings := []map[string]interface{}{
		{
			"name": "frete",
			"value": 1000,
		},
	}

	body := map[string]interface{} {
		"plan_id": 3,
		"customer": customer,
		"items": items,
		"shippings": shippings,
	}

	res, err := efi.UpdateSubscription(1, body) // no lugar do 1 coloque o subscription_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}