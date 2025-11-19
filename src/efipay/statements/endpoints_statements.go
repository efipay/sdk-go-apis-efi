package statements

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error)
	}
}

func (endpoints endpoints) ListStatementFiles(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/arquivos", "GEt", nil, body)
}

func (endpoints endpoints) GetStatementFile(params map[string]string) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/download/:nome_arquivo", "GEt", params, nil)
}

func (endpoints endpoints) ListStatementRecurrences(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/agendamentos", "GEt", nil, body)
}

func (endpoints endpoints) CreateStatementRecurrency(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/agendar", "POST", nil, body)
}

func (endpoints endpoints) UpdateStatementRecurrency(params map[string]string, body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/agendar/:identificador", "PATCH", params, body)
}

func (endpoints endpoints) CreateSftpKey(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/gerar-chaves", "POST", nil, body)
}
