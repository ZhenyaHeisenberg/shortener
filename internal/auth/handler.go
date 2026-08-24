package auth

import (
	"net/http"
	"project/configs"
	"project/pkg/request"
	"project/pkg/responce"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		_, err = handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			responce.Json(w, err.Error(), 400)
			return 
		}

		responce.Json(w, "Success", 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		_, err = handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			responce.Json(w, err.Error(), 400)
			return
		}
		responce.Json(w, "Success", 201)
	}
}
