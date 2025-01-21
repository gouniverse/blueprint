package middlewares

import (
	"net/http"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/router"
)

// NewAdminMiddleware checks if the user is an administrator or superuser
// before allowing access to the protected route.
//
// Business logic:
//  1. user must be authenticated
//  2. user must be active
//  3. user must be registered
//  4. user must be an admin or superuser
func NewAdminMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "Admin Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Admin validation logic here. CHange with your own

				authUser := helpers.GetAuthUser(r)

				// Check if user is authenticated? No => redirect to login
				if authUser == nil {
					returnURL := links.URL(r.URL.Path, map[string]string{})
					loginURL := links.NewAuthLinks().Login(returnURL)
					helpers.ToFlashError(w, r, "You must be logged in to access this page", loginURL, 15)
					return
				}

				if !authUser.IsRegistrationCompleted() {
					registerURL := links.NewAuthLinks().Register(map[string]string{})
					helpers.ToFlashInfo(w, r, "Please complete your registration to continue", registerURL, 15)
					return
				}

				// Check if user is active? No => redirect to website home
				if !authUser.IsActive() {
					homeURL := links.NewWebsiteLinks().Home()
					helpers.ToFlash(w, r, "error", "Your account is not active", homeURL, 15)
					return
				}

				// Check if user is an admin? No => redirect to website home
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
