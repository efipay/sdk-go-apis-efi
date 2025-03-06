package efipay

type Efipay struct {
	endpoints
}

func NewEfiPay(configs map[string]interface{}) *Efipay {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)
	//partner_token := configs["partner_token"].(string)

	requester := newRequester(clientID, clientSecret, sandbox, timeout)
	efi := Efipay{}
	efi.requester = *requester
	return &efi
}
