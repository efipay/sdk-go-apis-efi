package payments

type pix struct {
	endpoints
}

func NewEfiPay(configs map[string]interface{}) *pix {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	CA := configs["CA"].(string)
	Key := configs["Key"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requester := newRequester(clientID, clientSecret,CA, Key, sandbox, timeout)
	efi := pix{}
	efi.requester = *requester
	return &efi
}