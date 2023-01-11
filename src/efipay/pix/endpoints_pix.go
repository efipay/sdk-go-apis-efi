package pix

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error)
	}
}

func (endpoints endpoints) CreateImmediateCharge(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/cob", "POST", nil, body)
}


func (endpoints endpoints) CreateCharge(txid string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cob/:txid", "PUT", params, body)
}

func (endpoints endpoints) UpdateCharge(txid string) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cob/:txid", "PATCH", params, nil)
}

func (endpoints endpoints) DetailCharge(txid string) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cob/:txid", "GET", params, nil)
}

func (endpoints endpoints) ListCharges(inicio string, fim string) (string, error) {
	params := map[string]string{ 
		"inicio": (inicio),
		"fim": (fim),
	}
	return endpoints.requester.request("/v2/cob?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (endpoints endpoints) PixDevolution(e2eid string, id string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		 "e2eid": (e2eid),
		 "id": (id), }
	return endpoints.requester.request("/v2/pix/:e2eid/devolucao/:id", "PUT", params, body)
}

func (endpoints endpoints) PixDetailDevolution(e2eid string, id string) (string, error) {
	params := map[string]string{
		 "e2eid": (e2eid),
		 "id": (id), }
	return endpoints.requester.request("/v2/pix/:e2eid/devolucao/:id", "GET", params, nil)
}

func (endpoints endpoints) PixSend(idEnvio string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "idEnvio": (idEnvio) }
	return endpoints.requester.request("/v2/gn/pix/:idEnvio", "PUT", params, body)
}

func (endpoints endpoints) PixSendList(e2eid string) (string, error) {
	params := map[string]string{ "e2eid": (e2eid) }
	return endpoints.requester.request("/v2/pix/:e2eid", "GET", params, nil)
}

func (endpoints endpoints) PixSendDetail(e2eid string) (string, error) {
	params := map[string]string{ "e2eid": (e2eid) }
	return endpoints.requester.request("/v2/gn/pix/enviados/:e2eid", "GET", params, nil)
}

func (endpoints endpoints) PixCreateLocation(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/loc", "POST", nil, body)
}

func (endpoints endpoints) PixUnlinkTxidLocation(id string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/loc/:id/txid", "DELETE", params, body)
}

func (endpoints endpoints) PixDetailLocation(id string) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/loc/:id", "GET", params, nil)
}

func (endpoints endpoints) PixLocationList(inicio string, fim string) (string, error) {
	params := map[string]string{ 
		"inicio": (inicio),
		"fim": (fim),
	}
	return endpoints.requester.request("/v2/loc?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (endpoints endpoints) PixGenerateQRCode(id string) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/loc/:id/qrcode", "GET", params, nil)
}

func (endpoints endpoints) GetAccountBalance(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/saldo", "GET", nil, body)
}

func (endpoints endpoints) ListAccountConfig(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/config", "GET", nil, body)
}

func (endpoints endpoints) UpdateAccountConfig(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/config", "PUT", nil, body)
}

func (endpoints endpoints) PixCreateEvp(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/evp", "POST", nil, body)
}

func (endpoints endpoints) PixListEvp(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/evp", "GET", nil, body)
}

func (endpoints endpoints) PixDeleteEvp(key string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "key": (key) }
	return endpoints.requester.request("/v2/gn/evp/:key", "DELETE", params, body)
}

func (endpoints endpoints) PixConfigWebhook(chave string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "chave": (chave) }
	return endpoints.requester.request("/v2/webhook/:chave", "PUT", params, body)
}

func (endpoints endpoints) PixDeleteWebhook(chave string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "chave": (chave) }
	return endpoints.requester.request("/v2/webhook/:chave", "DELETE", params, body)
}

func (endpoints endpoints) PixDetailWebhook(chave string) (string, error) {
	params := map[string]string{ "chave": (chave) }
	return endpoints.requester.request("/v2/webhook/:chave", "GET", params, nil)
}

func (endpoints endpoints) PixListWebhooks(inicio string, fim string) (string, error) {
	params := map[string]string{ 
		"inicio": (inicio),
		"fim": (fim),
	}
	return endpoints.requester.request("/v2/webhook?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (endpoints endpoints) CreateDueCharge(txid string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cobv/:txid", "PUT", params, body)
}

func (endpoints endpoints) PixUpdateDueCharge(txid string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cobv/:txid", "PATCH", params, body)
}

func (endpoints endpoints) DetailDueCharge(txid string) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cobv/:txid", "GET", params, nil)
}

func (endpoints endpoints) PixListDueCharges(params map[string]string) (string, error) {
	return endpoints.requester.request("/v2/cobv", "GET", params, nil)
}

func (endpoints endpoints) CreateReport(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/relatorios/extrato-conciliacao", "POST", nil, body)
}

func (endpoints endpoints) DetailReport(id string) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/gn/relatorios/:id", "GET", params, nil)
}

func (endpoints endpoints) PixReceivedList(params map[string]string) (string, error) {
	return endpoints.requester.request("/v2/pix", "GET", params, nil)
}

func (endpoints endpoints) PixDetailReceived(e2eid string) (string, error) {
	params := map[string]string{ "e2eid": (e2eid) }
	return endpoints.requester.request("/v2/pix/:e2eid", "GET", params, nil)
}

func (endpoints endpoints) PixSplitDetailCharge(txid string) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/gn/split/cob/:txid", "GET", params, nil)
}

func (endpoints endpoints) PixSplitLinkCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{ 
		"txid": (txid),
		"splitConfigId": (splitConfigId),
	}
	return endpoints.requester.request("/v2/gn/split/cob/:txid/vinculo/:splitConfigId", "PUT", params, nil)
}

func (endpoints endpoints) PixSplitUnlinkCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{ 
		"txid": (txid),
		"splitConfigId": (splitConfigId),
	}
	return endpoints.requester.request("/v2/gn/split/cob/:txid/vinculo/:splitConfigId", "DELETE", params, nil)
}

func (endpoints endpoints) PixSplitDetailDueCharge(txid string) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/gn/split/cobv/:txid", "GET", params, nil)
}

func (endpoints endpoints) PixSplitLinkDueCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{ 
		"txid": (txid),
		"splitConfigId": (splitConfigId),
	}
	return endpoints.requester.request("/v2/gn/split/cobv/:txid/vinculo/:splitConfigId", "PUT", params, nil)
}

func (endpoints endpoints) PixSplitUnlinkDueCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{ 
		"txid": (txid),
		"splitConfigId": (splitConfigId),
	}
	return endpoints.requester.request("/v2/gn/split/cobv/:txid/vinculo/:splitConfigId", "DELETE", params, nil)
}

func (endpoints endpoints) PixSplitConfig(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/split/config", "POST", nil, body)
}

func (endpoints endpoints) PixSplitConfigId(id string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/gn/split/config/:id", "PUT", params, body)
}

func (endpoints endpoints) PixSplitDetailConfig(id string) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/gn/split/config/:id", "GET", params, nil)
}