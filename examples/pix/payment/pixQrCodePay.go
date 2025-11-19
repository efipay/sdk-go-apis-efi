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
		"idEnvio": "",
	}

	body := map[string]interface{}{
		"pagador": map[string]interface{}{
			"chave":       "",
			"infoPagador": "Pagamento de QR Code via API Pix",
		},
		"pixCopiaECola": "00020101021226830014BR.GOV.BCB.PIX2561qrcodespix.sejaefi.com.br/v2 41e0badf811a4ce6ad8a80b306821fce5204000053000065802BR5905EFISA6008SAOPAULO60070503***61040000",
	}

	res, err := efi.PixQrCodePay(params, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
