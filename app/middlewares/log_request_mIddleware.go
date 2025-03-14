package middlewares

import (
	"net/http"
	"project/config"
	"strings"

	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
)

// LogRequestMiddleware logs every request to the database using the LogStore logger
// ==================================================================
// This is userful so that we can identify where all the visits
// come from and keep the application protected - i.e. bots,
// malicious spiders, DDOS, etc
// ==================================================================
// it is useful to detect spamming bots
func NewLogRequestMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "Log Request Middleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uri := r.RequestURI

				ip := utils.IP(r)

				method := r.Method

				config.LogStore.InfoWithContext(method+": "+uri, map[string]string{
					"host":           r.Host,
					"path":           strings.TrimLeft(r.URL.Path, "/"),
					"ip":             ip,
					"method":         method,
					"useragent":      r.Header.Get("User-Agent"),
					"acceptlanguage": r.Header.Get("Accept-Language"),
					"referer":        r.Header.Get("Referer"),
				})

				next.ServeHTTP(w, r)
			})
		},
	}

	return m
}
