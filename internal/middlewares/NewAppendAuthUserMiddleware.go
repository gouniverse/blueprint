package middlewares

import (
	"context"
	"log"
	"net/http"
	"project/config"

	"github.com/gouniverse/auth"
	"github.com/gouniverse/router"
	"github.com/gouniverse/sessionstore"
	"github.com/gouniverse/utils"
)

func NewAppendAuthUserMiddleware() router.Middleware {
	m := router.Middleware{
		Name:    "Append Authenticated User Middleware",
		Handler: appendAuthUserHandler,
	}
	return m
}

// AppendUserMiddleware adds the user to the context.
//
// 1. The middleware checks if the user session key exists in the incoming
//    request. If it does not exist, the request is passed to the `next` handler.
//
// 2. If the user session key exists, it retrieves the user ID from the session
//    store using the session key.
//
// 2.1 If an error occurs while retrieving the user ID, the error is logged
// 	  and the request is passed to the `next` handler.
//
// 2.2 If the user ID is empty, the request is also passed to the `next` handler.
//
// 3. If the user ID is not empty, it queries the database to fetch the user.
//
// 3.1. If an error occurs during the user query, the error is logged
//		and the request is passed to the `next` handler.
//
// 4. If the user details are successfully fetched, it creates a new context
// 	  with the authenticated user value using `context.WithValue`.
//    The `config.AuthenticatedUserKey{}` is used as the key to store the user
//    value in the context.
//
// 6. The modified request with the new context is then passed to the `next` handler.
//
// Params:
//  - next http.Handler. The `next` handler is the next handler in the middleware chain.
// Returns
// - an http.Handler which represents the modified handler with the user.

func appendAuthUserHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userSessionKey := sessionKey(r)

		if userSessionKey == "" {
			next.ServeHTTP(w, r)
			return
		}

		userID, err := config.SessionStore.Get(userSessionKey, "", sessionstore.SessionOptions{
			IPAddress: utils.IP(r),
			UserAgent: r.UserAgent(),
		})

		if err != nil {
			log.Println(err.Error())
			next.ServeHTTP(w, r)
			return
		}

		if userID == "" {
			next.ServeHTTP(w, r)
			return
		}

		user, err := config.UserStore.UserFindByID(userID)

		if err != nil {
			log.Println(err.Error())
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), config.AuthenticatedUserKey{}, user)

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
