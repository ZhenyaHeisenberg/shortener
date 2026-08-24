package middleware

import "net/http"

type MiddleWare func(http.Handler) http.Handler

func Chain(middlewares ...MiddleWare) MiddleWare {
	return func(next http.Handler) http.Handler {
		for _, middleware := range middlewares {
			next = middleware(next)
		}
		return next
	}
}