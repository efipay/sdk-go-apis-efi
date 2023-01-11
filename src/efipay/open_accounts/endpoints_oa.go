package open_accounts


type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}, headers map[string]string) (string, error)
	}
}

func (endpoints endpoints) GetAccountCertificate() (string, error) {
	return endpoints.requester.request("/cadastro/conta-simplificada/:identificador/certificado", "GET", nil, nil, nil)
}

func (endpoints endpoints) GetAccountCredentials() (string, error) {
	return endpoints.requester.request("/cadastro/conta-simplificada/:identificador/credenciais", "GET", nil, nil, nil)
}

func (endpoints endpoints) CreateAccount(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/cadastro/conta-simplificada", "POST", nil, body, nil)
}

func (endpoints endpoints) AccountConfigWebhook(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/cadastro/webhook", "POST", nil, body, nil)
}

func (endpoints endpoints) AccountDeleteWebhook(params map[string]string) (string, error) {
	return endpoints.requester.request("/cadastro/webhook/:identificadorWebhook", "DELETE", params, nil, nil)
}

func (endpoints endpoints) AccountDetailWebhook(params map[string]string) (string, error) {
	return endpoints.requester.request("/cadastro/webhook/:identificadorWebhook", "GET", params, nil, nil)
}

func (endpoints endpoints) AccountListWebhook() (string, error) {
	return endpoints.requester.request("/cadastro/webhooks", "GET", nil, nil, nil)
}
