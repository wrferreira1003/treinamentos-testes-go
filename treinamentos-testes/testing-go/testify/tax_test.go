package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculateTax(1000.0)
	assert.Equal(t, 10.0, tax) // 10.0 == tax

	tax = CalculateTax(0)
	assert.Equal(t, 5.0, tax) // 5.0 == tax

	tax = CalculateTax(20000.0)
	assert.Equal(t, 20.0, tax) // 20.0 == tax

}

// Mocks no testes
// Testes de unidade - seria os testes que testam uma unidade do seu código
// Testes de integração - seria os testes que testam a integração de duas unidades
// Testes de aceitação - seria os testes que testam a aceitação do usuário
// Testes de sistema - seria os testes que testam o sistema como um todo
