package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wrferreira1003/api-server-go/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Teste de criação de usuário
func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	// Migrate the schema
	db.AutoMigrate(&entity.User{})

	// criando o usuário
	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user) // criando o usuário no banco de dados
	assert.NoError(t, err)    // verificando se não houve erros
	assert.NotNil(t, user.ID) // verificando se o usuário foi criado

	// buscando o usuário no banco de dados
	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

// Teste de busca de usuário por email
func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user) // criando o usuário no banco de dados
	assert.NoError(t, err)    // verificando se não houve erros
	assert.NotNil(t, user.ID) // verificando se o usuário foi criado

	// buscando o usuário no banco de dados
	userFound, err := userDB.FindByEmail(user.Email) // buscando o usuário por email
	assert.NoError(t, err)                           // verificando se não houve erros
	assert.NotNil(t, userFound)                      // verificando se o usuário foi encontrado
	assert.Equal(t, user.ID, userFound.ID)           // verificando se o ID do usuário é o mesmo
	assert.Equal(t, user.Name, userFound.Name)       // verificando se o nome do usuário é o mesmo
	assert.Equal(t, user.Email, userFound.Email)     // verificando se o email do usuário é o mesmo
	assert.NotNil(t, userFound.Password)             // verificando se a senha do usuário não é nula
}
