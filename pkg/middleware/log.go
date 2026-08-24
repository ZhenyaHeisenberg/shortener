package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     200,
		}
		next.ServeHTTP(wrapper, r)
		log.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(t))
	})
}
