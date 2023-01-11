package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	paymentToken := "428d7f3b2dc49117552ace464078557832c4ae4f";

	customer := map[string]interface{}{
		"name": "Gorbadoc Oldbuck",
		"cpf": "04267484171",
		"phone_number": "51944916523",
		"email": "gorb.oldbuck@gerencianet.com.br",
		"birth": "1977-01-15",
	}

	billingAddress := map[string]interface{} {
		"street": "Av JK",
		"number": 909,
		"neighborhood": "Bauxita",
		"zipcode": "35400000",
		"city": "Ouro Preto",
		"state": "MG",
	}

	body := map[string]interface{} {
		"payment": map[string]interface{} {
			"credit_card": map[string]interface{} {
			  "installments": 1,
			  "billing_address": billingAddress,
			  "payment_token": paymentToken,
			  "customer": customer,
			},
		},
		"items": []map[string]interface{}{
			{
				"name": "Product 1",
				"value": 1000,
				"amount": 2,
			},
		},
		"shippings": []map[string]interface{} {
			{
				"name": "Default Shipping Cost",
				"value": 100,
			},
		},
	}

	res, err := efi.CreateOneStepCharge(body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}