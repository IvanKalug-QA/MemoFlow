package main

import (
	"memoflow/configs"
	"memoflow/internal/auth"
	"memoflow/internal/memo"
	"memoflow/internal/stat"
	"memoflow/internal/user"
	"memoflow/pkg/db"
	"memoflow/pkg/middleware"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    config.Port.Name,
		Handler: stack(router),
	}

	// Repository
	memoRepository := memo.NewMemoRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	memo.NewMemoHandler(router, memo.MemoHandlerDeps{
		MemoResository: memoRepository,
		StatRepository: statRepository,
		Config:         config,
	})

	server.ListenAndServe()
}
