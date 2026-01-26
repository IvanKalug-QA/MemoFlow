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
}

type AuthHandler struct {
	*configs.Config
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
		fmt.Println("Register")
		var payload, err = req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(payload)
		data := RegisterResponse{Token: a.Auth.Secret}
		res.Json(w, data, http.StatusCreated)
	}
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{Config: deps.Config}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}
