package entity

import (
	"github.com/wrferreira1003/api-server-go/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       utils.ID
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-" is used to hide the field from the JSON output
}

// Todo usuario que for criado precisa passar pela função NewUser
func NewUser(name, email, password string) (*User, error) {
	// Hash da senha
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Criando o usuario
	return &User{
		ID:       utils.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

// Validando o password
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
