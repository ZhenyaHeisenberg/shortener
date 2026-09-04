package main

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/internal/auth"
	"project/internal/link"
	"project/internal/stat"
	"project/internal/user"
	"project/pkg/db"
	"project/pkg/middleware"
)

type Responce struct {
	StatusCode int
}

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf) // db
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		StatRepository: statRepository,
		Config:         conf,
	})

	// MiddleWares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Сервер запущен на порте 8081")
	server.ListenAndServe()
}
