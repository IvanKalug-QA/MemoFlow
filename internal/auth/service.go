package auth

import (
	"errors"
	"memoflow/internal/user"
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
	user := &user.User{
		Username: username,
		Email:    email,
		Password: "",
	}
	_, err := a.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
