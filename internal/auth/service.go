package auth

import (
	"errors"
	"memoflow/internal/user"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

type AuthService struct {
	UserRepository *user.UserRepository
}

func (a *AuthService) Register(email, password, username string) (string, error) {
	existedUser, _ := a.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	_, err = a.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
