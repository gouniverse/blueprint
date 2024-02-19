package middlewares

import (
	"net/http"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func NewAdminMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "Admin Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Admin validation logic here. CHange with your own

				authUser := helpers.GetAuthUser(r)

				if authUser == nil {
					returnURL := links.URL(r.URL.Path, map[string]string{})
					loginURL := links.NewAuthLinks().Login(returnURL)
					helpers.ToFlash(w, r, "error", "You must be logged in to access this page", loginURL, 15)
					return
				}

				if !authUser.IsActive() {
					homeURL := links.NewWebsiteLinks().Home()
					helpers.ToFlash(w, r, "error", "You must activate your account before you can access this page", homeURL, 15)
					return
				}

				if !authUser.IsAdministrator() && !authUser.IsSuperuser() {
					homeURL := links.NewWebsiteLinks().Home()
					helpers.ToFlash(w, r, "error", "You must be an administrator to access this page", homeURL, 15)
					return
				}

				next.ServeHTTP(w, r)
			})
		},
	}
	return m
}
