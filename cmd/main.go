package main

import (
	"fmt"
	"log"
	"net/http"
	"project/configs"
	"project/internal/auth"
	"project/internal/link"
	"project/internal/stat"
	"project/internal/user"
	"project/pkg/db"
	"project/pkg/event"
	"project/pkg/middleware"
)

type Responce struct {
	StatusCode int
}

func main() {
	conf := configs.LoadConfig()
	if conf == nil {
		log.Fatalln("config is nil")
		return
	}
	db := db.NewDb(conf) // db
		if conf == nil {
		log.Fatalln("db is nil")
		return
	}
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	statServise := stat.NewStatService(stat.StatServiceDeps{
		EventBus: eventBus,
		StatRepository: statRepository,
	})

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config: conf,
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

	go statServise.AddClick()

	fmt.Println("Сервер запущен на порте 8081")
	server.ListenAndServe()
}
