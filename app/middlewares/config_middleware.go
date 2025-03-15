package middlewares

import (
	"net/http"
	"project/config_v2"

	"github.com/gouniverse/router"
	"github.com/mingrammer/cfmt"
)

var ERROR_CONFIG_NOT_FOUND = `sorry, there was a config error`

// NewConfigMiddleware attaches the config_v2 configuration to the request context.
// This allows controllers and other middleware to access the configuration
// using config_v2.FromContext(r.Context()).
//
// Business logic:
//  1. Get the config instance
//  2. Add it to the request context
//  3. Pass the modified request to the next handler
func NewConfigMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "Config Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cfg, err := config_v2.New()

				if err != nil {
					cfmt.Error("config_middleware: ", "error: ", err.Error())
					http.Error(w, ERROR_CONFIG_NOT_FOUND, http.StatusInternalServerError)
					return
				}

				// Add the config to the request context
				ctx := config_v2.ToContext(r.Context(), cfg)

				// Create a new request with the updated context
				r = r.WithContext(ctx)

				// Call the next handler with the updated request
				next.ServeHTTP(w, r)
			})
		},
	}
	return m
}
