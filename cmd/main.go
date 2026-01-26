package main

import (
	"memoflow/configs"
	"memoflow/internal/auth"
	"memoflow/pkg/db"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	_ = db.NewDb(config)
	router := http.NewServeMux()
	server := http.Server{
		Addr:    config.Port.Name,
		Handler: router,
	}

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: config})

	server.ListenAndServe()
}
