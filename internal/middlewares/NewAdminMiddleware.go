package middlewares

import (
	"net/http"

	"github.com/gouniverse/router"
)

func NewAdminMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "Admin Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Admin validation logic here
				next.ServeHTTP(w, r)
			})
		},
	}
	return m
}
