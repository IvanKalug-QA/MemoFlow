package auth

import (
	"errors"
	"memoflow/internal/user"
	"memoflow/pkg/di"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

type AuthService struct {
	UserRepository di.IUserRepository
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

func (a *AuthService) Login(email, password string) error {
	existedUser, _ := a.UserRepository.FindByEmail(email)
	if existedUser == nil {
		return errors.New(ErrUserNotExists)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
