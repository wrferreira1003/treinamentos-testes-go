package database

import (
	"github.com/wrferreira1003/api-server-go/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

// Cria um novo produto
func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

// Implementa os métodos da interface UserInterface
func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

// Implementa preenchendo com os dados do usuario
func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
