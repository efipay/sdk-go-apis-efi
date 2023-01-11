package configs

var Credentials = map[string]interface{} {
	"client_id": "",
	"client_secret": "",
	"sandbox": false,
	"timeout": 20,
	"CA" : "", //caminho da chave publica da gerencianet
	"Key" : "", //caminho da chave privada da sua conta Gerencianet
}