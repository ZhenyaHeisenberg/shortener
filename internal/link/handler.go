package link

import (
	"fmt"
	"net/http"
	"project/pkg/request"
	"project/pkg/responce"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("POST /link/createlink", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) GoTo() http.HandlerFunc { // GET
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			responce.Json(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc { // POST
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		responce.Json(w, createdLink, 201)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc { // PATCH
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc { // DELETE
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
