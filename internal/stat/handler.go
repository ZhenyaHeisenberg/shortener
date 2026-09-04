package stat

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/pkg/middleware"
	"project/pkg/responce"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.URL)

		from, err := time.Parse("2006-1-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", 400)
			return
		}

		to, err := time.Parse("2006-1-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", 400)
			return
		}

		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			http.Error(w, "Invalid by param", 400)
			return
		}

		fmt.Println(from, to, by)
		responce.Json(w, "Все круто", 200)
	}
}
