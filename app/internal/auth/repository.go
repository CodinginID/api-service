package auth

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUsersAfterDate(date string) ([]User, error)
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

func (r *authRepository) GetUserByUsername(username string) (*User, error) {
	var user User
	fmt.Println("Fetching user by username:", username)
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) GetUsersAfterDate(date string) ([]User, error) {
	var users []User
	log.Default().Println("Fetching users after date:", date)
	if err := r.db.Where("created_at > ?", date).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
