package link

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/pkg/request"
	"project/pkg/responce"
)

type LinkHandlerDeps struct {
	*configs.Config
}

type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{}

	router.HandleFunc("GET /link/{hash}", handler.GoTo())
	router.HandleFunc("POST /link/createlink", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) GoTo() http.HandlerFunc { // GET
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LinkHandler](&w, r)
		if err != nil {
			responce.Json(w, err, 400)
			return
		}
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc { // POST
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LinkHandler](&w, r)
		if err != nil {
			responce.Json(w, err, 400)
			return
		}
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc { // PATCH
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LinkHandler](&w, r)
		if err != nil {
			responce.Json(w, err, 400)
			return
		}
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc { // DELETE
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
