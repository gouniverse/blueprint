package middlewares

import (
	"net/http"
	"project/config"

	"github.com/gouniverse/cmsstore"
	"github.com/mingrammer/cfmt"
)

func CmsAddMiddlewares() {
	if config.CmsStoreUsed {
		return
	}

	if config.CmsStore == nil {
		return
	}

	helloMiddleware := cmsstore.Middleware().
		SetIdentifier("HelloMiddleware").
		SetName("HelloMiddleware").
		SetType(cmsstore.MIDDLEWARE_TYPE_BEFORE).
		SetHandler(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cfmt.Infoln("Hello from Middleware")
				next.ServeHTTP(w, r)
			})
		})
	afterMiddleware := cmsstore.Middleware().
		SetIdentifier("CmsLayoutMiddleware").
		SetName("Cms Layout Middleware").
		SetType(cmsstore.MIDDLEWARE_TYPE_AFTER).
		SetHandler(NewCmsLayoutMiddleware().Handler)

	config.CmsStore.AddMiddleware(helloMiddleware)
	config.CmsStore.AddMiddleware(afterMiddleware)
}
