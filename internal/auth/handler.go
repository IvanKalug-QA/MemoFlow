package auth

import (
	"fmt"
	"memoflow/configs"
	"memoflow/pkg/req"
	"memoflow/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")
		var payload, err = req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(payload)
		data := LoginResponse{Token: a.Auth.Secret}
		res.Json(w, data, http.StatusOK)
	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body, err = req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		a.AuthService.Register(body.Email, body.Password, body.Username)
	}
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{Config: deps.Config, AuthService: deps.AuthService}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}
