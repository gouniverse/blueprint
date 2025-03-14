package middlewares

import (
	"net/http"
	"project/internal/helpers"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/router"
)

// NewUserMiddleware checks if the user is authenticated and active
// before allowing access to the protected route.
//
// Business logic:
//  1. user must be authenticated
//  2. user must be active
//  3. user must be registered
func NewUserMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "User Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// User validation logic here. Change with your own

				authUser := helpers.GetAuthUser(r)

				// Check if user is authenticated? No => redirect to login
				if authUser == nil {
					returnURL := links.URL(r.URL.Path, map[string]string{})
					loginURL := links.NewAuthLinks().Login(returnURL)
					helpers.ToFlashError(w, r, "Only authenticated users can access this page", loginURL, 15)
					return
				}

				// Check if user is active? No => redirect to website home
				if !authUser.IsActive() {
					homeURL := links.NewWebsiteLinks().Home()
					helpers.ToFlashError(w, r, "User account not active", homeURL, 15)
					return
				}

				// Check if user has completed registration? No => redirect to profile to complete registration
				notOnProfilePage := strings.Trim(r.URL.Path, "/") != strings.Trim(links.USER_PROFILE, "/") &&
					strings.Trim(r.URL.Path, "/") != strings.Trim(links.AUTH_REGISTER, "/")

				if !authUser.IsRegistrationCompleted() && notOnProfilePage {
					registerURL := links.NewAuthLinks().Register(map[string]string{})
					helpers.ToFlashInfo(w, r, "Please complete your registration to continue", registerURL, 15)
					return
				}

				next.ServeHTTP(w, r)
			})
		},
	}
	return m
}
