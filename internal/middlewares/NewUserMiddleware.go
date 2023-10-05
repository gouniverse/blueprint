package middlewares

import (
	"net/http"

	"github.com/gouniverse/router"
)

func NewUserMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "User Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// User validation logic here
				next.ServeHTTP(w, r)
			})
		},
	}
	return m
}
