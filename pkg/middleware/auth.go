package middleware

import (
	"context"
	"net/http"
	"project/configs"
	"project/pkg/jwt"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer") {
			writeUnauthorized(w)
			return 
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(conf.Auth.Secret).Parse(token)
		if !isValid || data == nil {
			writeUnauthorized(w)
			return 
		}
		r.Context()
		cxt := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(cxt)

		next.ServeHTTP(w, req)
	})
}
