package payments

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error)
	}
}

func (endpoints endpoints) PayDetailBarCode(params map[string]string) (string, error) {
	return endpoints.requester.request("/codBarras/:codBarras", "GET", params, nil)
}

func (endpoints endpoints) PayDetailPayment(params map[string]string)(string, error) {
	return endpoints.requester.request("/:idPagamento", "GET", params, nil)
}

func (endpoints endpoints) PayListPayments(params map[string]string) (string, error) {
	return endpoints.requester.request("/resumo", "GET", params, nil)
}

func (endpoints endpoints) PayRequestBarCode(params map[string]string, body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/codBarras/:codBarras", "POST", params, body)
}
