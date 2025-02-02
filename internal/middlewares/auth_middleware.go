package middlewares

import (
	"context"
	"log"
	"net/http"
	"project/config"

	"github.com/gouniverse/auth"
	"github.com/gouniverse/router"
)

func NewAuthMiddleware() router.Middleware {
	m := router.Middleware{
		Name:    "Auth Middleware",
		Handler: authHandler,
	}
	return m
}

// authHandler adds the user and session to the context.
//
//  1. Checks if the user session key exists in the incoming request.
//
//  2. Retrieves the session using the session key..
//
//  3. Checks the session is not expired.
//
//  4. Retrieves the user using the user ID from the session.
//
// Params:
//   - next http.Handler. The `next` handler is the next handler in the middleware chain.
//
// Returns
// - an http.Handler which represents the modified handler with the user.
func authHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionKey := authHandlerSessionKey(r)

		if sessionKey == "" {
			next.ServeHTTP(w, r)
			return
		}

		session, err := config.SessionStore.SessionFindByKey(sessionKey)

		if err != nil {
			config.Logger.Error("auth_middleware", "error", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		if session == nil {
			next.ServeHTTP(w, r)
			return
		}

		if session.IsExpired() {
			next.ServeHTTP(w, r)
			return
		}

		userID := session.GetUserID()

		if userID == "" {
			next.ServeHTTP(w, r)
			return
		}

		user, err := config.UserStore.UserFindByID(r.Context(), userID)

		if err != nil {
			config.Logger.Error("auth_middleware", "error", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		if user == nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), config.AuthenticatedUserContextKey{}, user)
		ctx = context.WithValue(ctx, config.AuthenticatedSessionContextKey{}, sessionKey)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// authHandlerSessionKey returns the session key from the incoming request.
func authHandlerSessionKey(r *http.Request) string {
	authTokenFromCookie, err := r.Cookie(auth.CookieName)

	if err != nil {
		if err != http.ErrNoCookie {
			log.Println(err.Error())
		}
	}

	if authTokenFromCookie == nil {
		return ""
	}

	sessionKey := authTokenFromCookie.Value

	return sessionKey
}
