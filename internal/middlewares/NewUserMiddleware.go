package middlewares

import (
	"net/http"
	"project/internal/helpers"
	"project/internal/links"
	"project/pkg/userstore"
	"strings"

	"github.com/gouniverse/router"
)

// NewUserMiddleware checks if the user is authenticated and active
// before allowing access to the protected route.
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
					helpers.ToFlash(w, r, "error", "Only authenticated users can access this page", loginURL, 15)
					return
				}

				// Check if user is active? No => redirect to website home
				if !authUser.IsActive() {
					homeURL := links.NewWebsiteLinks().Home()
					helpers.ToFlash(w, r, "error", "User account not active", homeURL, 15)
					return
				}

				// Check if user has completed registration? No => redirect to profile to complete registration
				notOnProfilePage := strings.Trim(r.URL.Path, "/") != strings.Trim(links.USER_PROFILE, "/")

				if isUserRegistrationIncomplete(authUser) && notOnProfilePage {
					homeURL := links.NewUserLinks().Profile()
					helpers.ToFlashInfo(w, r, "Please complete your registration to continue", homeURL, 15)
					return
				}

				next.ServeHTTP(w, r)
			})
		},
	}
	return m
}

// isUserRegistrationIncomplete checks if the user is incomplete.
//
// Registration is considered incomplete if the user's first name
// or last name is empty.
//
// Parameters:
// - authUser: a pointer to a userstore.User object representing the authenticated user.
//
// Returns:
// - bool: true if the user registration is incomplete, false otherwise.
func isUserRegistrationIncomplete(authUser *userstore.User) bool {
	return authUser.FirstName() == "" && authUser.LastName() == ""
}
