package main

import (
	"memoflow/configs"
	"memoflow/internal/auth"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	server := http.Server{
		Addr:    config.Port.Name,
		Handler: router,
	}

	auth.NewAuthHandler(router)

	server.ListenAndServe()
}
