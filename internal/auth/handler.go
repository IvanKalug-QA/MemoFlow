package auth

import (
	"memoflow/configs"
	"memoflow/pkg/jwt"
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
		var payload, err = req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		err = a.AuthService.Login(payload.Email, payload.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(a.Auth.Secret).Create(jwt.JWTData{Email: payload.Email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		loginResponse := LoginResponse{Token: token}
		res.Json(w, loginResponse, http.StatusOK)
	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body, err = req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		email, err := a.AuthService.Register(body.Email, body.Password, body.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token, err := jwt.NewJWT(a.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		registerResponse := RegisterResponse{Token: token}
		res.Json(w, registerResponse, http.StatusOK)
	}
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{Config: deps.Config, AuthService: deps.AuthService}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}
