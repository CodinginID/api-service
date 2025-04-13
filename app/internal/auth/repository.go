package auth

import (
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
