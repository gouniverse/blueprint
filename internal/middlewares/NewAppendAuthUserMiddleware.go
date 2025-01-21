package middlewares

import (
	"context"
	"log"
	"net/http"
	"project/config"

	"github.com/gouniverse/auth"
	"github.com/gouniverse/router"
)

func NewAppendAuthUserMiddleware() router.Middleware {
	m := router.Middleware{
		Name:    "Append Authenticated User Middleware",
		Handler: AppendUserHandler,
	}
	return m
}

// AppendUserHandler adds the user to the context.
//
//  1. The middleware checks if the user session key exists in the incoming
//     request. If it does not exist, the request is passed to the `next` handler.
//
//  2. If the user session key exists, it retrieves the user ID from the session
//     store using the session key.
//
// 2.1 If an error occurs while retrieving the user ID, the error is logged
//
//	and the request is passed to the `next` handler.
//
// 2.2 If the user ID is empty, the request is also passed to the `next` handler.
//
// 3. If the user ID is not empty, it queries the database to fetch the user.
//
// 3.1. If an error occurs during the user query, the error is logged
//
//		and the request is passed to the `next` handler.
//
//	 4. If the user details are successfully fetched, it creates a new context
//	    with the authenticated user value using `context.WithValue`.
//	    The `config.AuthenticatedUserContextKey{}` is used as the key to store the user
//	    value in the context.
//
// 6. The modified request with the new context is then passed to the `next` handler.
//
// Params:
//   - next http.Handler. The `next` handler is the next handler in the middleware chain.
//
// Returns
// - an http.Handler which represents the modified handler with the user.
func AppendUserHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userSessionKey := sessionKey(r)

		if userSessionKey == "" {
			next.ServeHTTP(w, r)
			return
		}

		session, err := config.SessionStore.SessionFindByKey(userSessionKey)

		if err != nil {
			config.Logger.Error("At appendUserHandler", "error", err.Error())
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
			config.Logger.Error("At appendUserHandler", "error", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), config.AuthenticatedUserContextKey{}, user)
		ctx = context.WithValue(ctx, config.AuthenticatedSessionContextKey{}, userSessionKey)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func sessionKey(r *http.Request) string {
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
