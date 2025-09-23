package open_finance

type Of struct {
	endpoints
}

func NewEfiPay(configs map[string]interface{}) *Of {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	CA := configs["CA"].(string)
	Key := configs["Key"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requester := newRequester(clientID, clientSecret, CA, Key, sandbox, timeout)
	efi := Of{}
	efi.requester = *requester
	return &efi
}
