package main

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/pkg/db"
	"project/internal/auth"
)

type Responce struct {
	StatusCode int
}

func main() {
	conf := configs.LoadConfig()
	db.NewDb(conf) // db
	router := http.NewServeMux()
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
