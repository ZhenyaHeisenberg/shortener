package main

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/internal/auth"
	"project/internal/link"
	"project/pkg/db"
)

type Responce struct {
	StatusCode int
}

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf) // db
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handlers
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Сервер запущен на порте 8081")
	server.ListenAndServe()
}
