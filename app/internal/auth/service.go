package auth

import (
	"errors"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, email, password string) error
	Login(email, password string) (string, error)
	GetUserDetailByUsername(username string) (*User, error)
	GetUsersAfterDate(date string) ([]User, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(username, email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Check if user already exists
	existingUser, err := s.repo.GetUserByEmail(email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}

	// Create new user

	user := &User{
		Username:  username,
		Email:     email,
		Password:  string(hashed),
		CreatedAt: time.Now(),
	}

	return s.repo.CreateUser(user)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func (s *authService) GetUserDetailByUsername(username string) (*User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) GetUsersAfterDate(date string) ([]User, error) {
	users, err := s.repo.GetUsersAfterDate(date)
	if err != nil {
		return nil, err
	}
	return users, nil
}
