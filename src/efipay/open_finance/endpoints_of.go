package open_finance

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}, headers map[string]string) (string, error)
	}
}

func (endpoints endpoints) OfListParticipants(nome string) (string, error) {
	params := map[string]string{ "nome": (nome) }
	return endpoints.requester.request("/participantes", "GET", params, nil, nil)
}

func (endpoints endpoints) OfConfigUpdate(body map[string]interface{})(string, error) {
	return endpoints.requester.request("/config", "PUT", nil, body, nil)
}

func (endpoints endpoints) OfConfigDetail() (string, error) {
	return endpoints.requester.request("/config", "GET", nil, nil, nil)
}

func (endpoints endpoints) OfDevolutionPix(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/devolucao/pagamento/pix", "POST", nil, body, nil)
}

func (endpoints endpoints) OfListPixPayment(inicio string, fim string) (string, error) {
	params := map[string]string{ 
		"inicio": (inicio),
		"fim": (fim),
	}
	return endpoints.requester.request("/pagamentos/pix", "GET", params, nil, nil)
}

func (endpoints endpoints) OfStartPixPayment(body map[string]interface{}, headers map[string]string) (string, error) {
	return endpoints.requester.request("/pagamentos/pix", "POST", nil, body, headers)
}

