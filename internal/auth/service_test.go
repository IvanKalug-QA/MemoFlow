package auth_test

import (
	"memoflow/internal/auth"
	"memoflow/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (r *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}
func (r *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@a.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register("a@a.ru", "1", "Vasya")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Email %s do not math %s", email, initialEmail)
	}
}
