package open_finance

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}, headers map[string]string) (string, error)
	}
}

func (endpoints endpoints) OfListParticipants(nome string) (string, error) {
	params := map[string]string{"nome": (nome)}
	return endpoints.requester.request("/participantes", "GET", params, nil, nil)
}

func (endpoints endpoints) OfConfigUpdate(body map[string]interface{}) (string, error) {
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
		"fim":    (fim),
	}
	return endpoints.requester.request("/pagamentos/pix", "GET", params, nil, nil)
}

func (endpoints endpoints) OfStartPixPayment(body map[string]interface{}, headers map[string]string) (string, error) {
	return endpoints.requester.request("/pagamentos/pix", "POST", nil, body, headers)
}

func (endpoints endpoints) OfStartRecurrencyPixPayment(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "POST", nil, body, nil)
}

func (endpoints endpoints) OfListRecurrencyPixPayment(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "GET", params, nil, nil)
}

func (endpoints endpoints) OfCancelRecurrencyPix(identificadorPagamento string) (string, error) {
	params := map[string]string{"identificadorPagamento": (identificadorPagamento)}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/cancelar", "PATCH", params, nil, nil)
}

func (endpoints endpoints) OfDevolutionRecurrencyPix(identificadorPagamento string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		"identificadorPagamento": (identificadorPagamento),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/devolver", "POST", nil, body, params)
}

func (endpoints endpoints) OfReplaceRecurrencyPixParcel(identificadorPagamento string, endToEndId string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		"identificadorPagamento": (identificadorPagamento),
		"endToEndId":             (endToEndId),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/substituir/:endToEndId", "PATCH", nil, body, params)
}

func (endpoints endpoints) OfCancelSchedulePix(identificadorPagamento string) (string, error) {
	params := map[string]string{"identificadorPagamento": (identificadorPagamento)}
	return endpoints.requester.request("/v1/pagamentos-agendados/pix/:identificadorPagamento/cancelar", "PATCH", params, nil, nil)
}

func (endpoints endpoints) OfListSchedulePixPayment(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "GET", params, nil, nil)
}

func (endpoints endpoints) OfStartSchedulePixPayment(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "POST", nil, body, nil)
}

func (endpoints endpoints) OfDevolutionSchedulePix(identificadorPagamento string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		"identificadorPagamento": (identificadorPagamento),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/devolver", "POST", nil, body, params)
}

func (endpoints endpoints) OfCreateBiometricEnrollment(body map[string]interface{}, headers map[string]string) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-biometria/vinculos", "POST", nil, body, headers)
}

func (endpoints endpoints) OfListBiometricEnrollment(requestParams map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-biometria/vinculos", "GET", nil, requestParams, nil)
}

func (endpoints endpoints) OfRevokeBiometricEnrollment(body map[string]interface{}, headers map[string]string) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-biometria/vinculos", "PATCH", nil, body, headers)
}

func (endpoints endpoints) OfCreateBiometricPixPayment(body map[string]interface{}, headers map[string]string) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-biometria/pix", "POST", nil, body, headers)
}

func (endpoints endpoints) OfListBiometricPixPayment(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return endpoints.requester.request("/v1/pagamentos-biometria/pix", "GET", params, nil, nil)
}

func (endpoints endpoints) OfCreateAutomaticEnrollment(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-automaticos/adesao", "POST", nil, body, nil)
}

func (endpoints endpoints) OfListAutomaticEnrollment(requestParams map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-automaticos/adesao", "GET", nil, requestParams, nil)
}

func (endpoints endpoints) OfUpdateAutomaticEnrollment(body map[string]interface{}, headers map[string]string) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-automaticos/adesao", "PATCH", nil, body, headers)
}

func (endpoints endpoints) OfCreateAutomaticPixPayment(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-automaticos/pix", "POST", nil, body, nil)
}

func (endpoints endpoints) OfListAutomaticPixPayment(requestParams map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-automaticos/pix", "GET", nil, requestParams, nil)
}

func (endpoints endpoints) OfCancelAutomaticPixPayment(requestParams map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-automaticos/pix", "PATCH", nil, requestParams, nil)
}
