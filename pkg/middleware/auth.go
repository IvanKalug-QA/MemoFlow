package middleware

import (
	"context"
	"memoflow/configs"
	"memoflow/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func WriteUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if !strings.HasPrefix(bearer, "Bearer ") {
			WriteUnauthed(w)
			return
		}
		token := strings.TrimPrefix(bearer, "Bearer ")
		is_valid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !is_valid {
			WriteUnauthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
