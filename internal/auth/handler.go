package auth

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}

func (ul *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")
	}
}

func (ul *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}

func NewAuthHandler(router *http.ServeMux) {
	handler := &AuthHandler{}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}
