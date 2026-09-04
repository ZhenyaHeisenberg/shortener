package link

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/pkg/event"
	"project/pkg/middleware"
	"project/pkg/request"
	"project/pkg/responce"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
	Config         *configs.Config
}

type LinkHandler struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		EventBus:       deps.EventBus,
	}

	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.Handle("GET /link", middleware.IsAuthed(handler.GetAll(), deps.Config))
	router.HandleFunc("POST /link/createlink", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())

}

func (handler *LinkHandler) GoTo() http.HandlerFunc { // GET
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		go handler.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: link.ID,
		})
		http.Redirect(w, r, link.Url, 307)
	}
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "invalid limit", 400)
			return
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "invalid offset", 400)
			return
		}

		links := handler.LinkRepository.GetAll(limit, offset)
		count := handler.LinkRepository.Count()
		resp := GetAllLinksResponse{
			Links: links,
			Count: int(count),
		}
		responce.Json(w, resp, 200)
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc { // POST
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		responce.Json(w, createdLink, 201)

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc { // PATCH
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println("Вытащилли Email:", email)
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if body.Hash != "" {

		}

		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		responce.Json(w, link, 201)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc { // DELETE
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		_, err = handler.LinkRepository.FindById(uint(id)) // Проверка наличия
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		err = handler.LinkRepository.Delete(uint(id)) // Удаление
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		responce.Json(w, nil, 204)
	}
}
