package main

import (
	"fmt"
	"github.com/efipay/sdk-go-apis-efi/src/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

   paymentToken := "Insira_aqui_seu_paymentToken";

  repassOne := map[string]interface{}{
    "payee_code": "informe_payee_code_conta", // identificador da conta Gerencianet (repasse 1)
    "percentage": 2500, // porcentagem de repasse (2500 = 25%)
  }

  repassTwo := map[string]interface{}{
    "payee_code": "informe_payee_code_conta", // identificador da conta Gerencianet (repasse 2)
    "percentage": 1500, // porcentagem de repasse (1500 = 15%)
  }
  
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
        "value": 500,
        "amount": 1,
        "marketplace": map[string]interface{}{
          "repasses": []map[string]interface{}{
          repassOne,
          repassTwo,
          },
        },
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
