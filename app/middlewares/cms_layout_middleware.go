package middlewares

import (
	"net/http"
	"net/http/httptest"
	"project/app/layouts"

	"github.com/gouniverse/cmsstore"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/samber/lo"
)

// NewCmsLayoutMiddleware is a middleware that is specific to the CMS.
//
// The middleware can be attached to the page via the CMS admin interface
// as an "after" middleware.
//
// The middleware is responsible for rendering the CMS pages. It wraps the
// original page content with the user dashboard layout, allowing the CMS
// pages to become one whole with the overall portal, which includes the
// navigation header (with login and logout links).
//
// It uses the "page" context value to transfer the page data (i.e. title,
// meta keywords, description) from the CMS frontend to the layout.
func NewCmsLayoutMiddleware() router.Middleware {
	m := router.Middleware{
		Name: "CmsLayoutMiddleware",
		Handler: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				page, ok := r.Context().Value("page").(cmsstore.PageInterface)
				if !ok {
					page = nil
				}

				title := lo.
					If(page == nil, "").
					ElseF(
						func() string {
							return page.Title()
						})

				rec := httptest.NewRecorder()
				next.ServeHTTP(rec, r)
				finalContent := rec.Body.String()

				fullPage := layouts.NewUserLayout(r, layouts.Options{
					Title:      title,
					Content:    hb.Raw(finalContent),
					ScriptURLs: []string{},
					Styles:     []string{},
				}).ToHTML()

				w.Write([]byte(fullPage))
				return
			})
		},
	}

	return m
}
