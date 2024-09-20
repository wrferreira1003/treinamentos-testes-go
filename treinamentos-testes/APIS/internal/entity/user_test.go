package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Garantir que o usuario seja criado corretamente

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@example.com", "123456")
	assert.NoError(t, err)                          // Garantir que não tenha erro
	assert.NotNil(t, user)                          // Garantir que o usuario não seja nulo
	assert.NotEmpty(t, user.ID)                     // Garantir que o ID não esteja vazio
	assert.Equal(t, "John Doe", user.Name)          // Garantir que o nome seja igual ao informado
	assert.Equal(t, "john@example.com", user.Email) // Garantir que o email seja igual ao informado
	assert.NotEmpty(t, user.Password)               // Garantir que a senha não esteja vazia
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john@example.com", "123456")
	assert.NoError(t, err)                            // Garantir que não tenha erro
	assert.True(t, user.ValidatePassword("123456"))   // Garantir que a senha seja valida
	assert.False(t, user.ValidatePassword("1234567")) // Garantir que a senha seja invalida
	assert.NotEqual(t, "123456", user.Password)       // Garantir que a senha não seja igual a informada
}
