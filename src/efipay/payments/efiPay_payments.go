package payments

type Payment struct {
	endpoints
}

func NewEfiPay(configs map[string]interface{}) *Payment {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	CA := configs["CA"].(string)
	Key := configs["Key"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requester := newRequester(clientID, clientSecret, CA, Key, sandbox, timeout)
	efi := Payment{}
	efi.requester = *requester
	return &efi
}
