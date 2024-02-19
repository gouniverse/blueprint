package middlewares

import (
	"net/http"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func NewUserMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "User Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// User validation logic here. Change with your own

				authUser := helpers.GetAuthUser(r)

				if authUser == nil {
					returnURL := links.URL(r.URL.Path, map[string]string{})
					loginURL := links.NewAuthLinks().Login(returnURL)
					helpers.ToFlash(w, r, "error", "Only authenticated users can access this page", loginURL, 15)
					return
				}

				if !authUser.IsActive() {
					homeURL := links.NewWebsiteLinks().Home()
					helpers.ToFlash(w, r, "error", "User account not active", homeURL, 15)
					return
				}

				next.ServeHTTP(w, r)
			})
		},
	}
	return m
}
