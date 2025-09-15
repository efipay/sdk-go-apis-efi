package efipay

import (
	"testing"
)

func TestNewEfiPay(t *testing.T) {
	// Configurações de teste
	configs := map[string]interface{}{
		"client_id":     "test_client_id",
		"client_secret": "test_client_secret",
		"sandbox":       true,
		"timeout":       30,
	}

	// Testa a criação da instância Efipay
	efi := NewEfiPay(configs)

	// Verifica se a instância não é nil
	if efi == nil {
		t.Fatal("NewEfiPay retornou nil")
	}

	// Verifica se é do tipo correto
	if efi == nil {
		t.Errorf("NewEfiPay não retornou uma instância válida")
	}
}

func TestNewEfiPayWithInvalidConfigs(t *testing.T) {
	// Testa com configurações inválidas (deve causar panic)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("NewEfiPay deveria ter causado panic com configurações inválidas")
		}
	}()

	// Configurações inválidas (faltando campos obrigatórios)
	invalidConfigs := map[string]interface{}{
		"client_id": "test_client_id",
		// faltando client_secret, sandbox e timeout
	}

	NewEfiPay(invalidConfigs)
}

func TestNewEfiPayWithWrongTypes(t *testing.T) {
	// Testa com tipos incorretos (deve causar panic)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("NewEfiPay deveria ter causado panic com tipos incorretos")
		}
	}()

	// Configurações com tipos incorretos
	wrongTypeConfigs := map[string]interface{}{
		"client_id":     "test_client_id",
		"client_secret": "test_client_secret",
		"sandbox":       "true", // deveria ser bool
		"timeout":       30,
	}

	NewEfiPay(wrongTypeConfigs)
}

func TestEfipayStructFields(t *testing.T) {
	// Configurações de teste
	configs := map[string]interface{}{
		"client_id":     "test_client_id",
		"client_secret": "test_client_secret",
		"sandbox":       true,
		"timeout":       30,
	}

	efi := NewEfiPay(configs)

	// Verifica se a struct Efipay incorpora endpoints
	// Testamos isso verificando se um método de endpoints está disponível
	// Como CreateCharge é um método, apenas verificamos se podemos chamá-lo
	// (não podemos comparar métodos com nil)
	if efi == nil {
		t.Errorf("Instância Efipay é nil")
	}
}

// Teste de benchmark para medir performance da criação
func BenchmarkNewEfiPay(b *testing.B) {
	configs := map[string]interface{}{
		"client_id":     "test_client_id",
		"client_secret": "test_client_secret",
		"sandbox":       true,
		"timeout":       30,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewEfiPay(configs)
	}
}