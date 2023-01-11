package efipay

type efipay struct {
	endpoints
}

func NewEfiPay(configs map[string]interface{}) *efipay {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requester := newRequester(clientID, clientSecret, sandbox, timeout)
	efi := efipay{}
	efi.requester = *requester
	return &efi
}
