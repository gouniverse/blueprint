package admin

import (
	"net/http"
	"project/internal/links"

	"github.com/gouniverse/router"
)

func BlogRoutes() []router.RouteInterface {
	handler := func(w http.ResponseWriter, r *http.Request) string {
		return NewBlogPostManagerController().Handler(w, r)
	}

	return []router.RouteInterface{
		&router.Route{
			Name:        "Admin > Blog > Post Create",
			Path:        links.ADMIN_BLOG_POST_CREATE,
			HTMLHandler: NewPostCreateController().Handler,
		},
		&router.Route{
			Name:        "Admin > Blog > Post Delete",
			Path:        links.ADMIN_BLOG_POST_DELETE,
			HTMLHandler: NewPostDeleteController().Handler,
		},
		// {
		// 	Name:    "Admin > Blog > Post Details",
		// 	Path:    links.ADMIN_BLOG_POST_VIEW,
		// 	Handler: adminBlog.NewPostViewController().Handler,
		// },
		&router.Route{
			Name:        "Admin > Blog > Post Manager",
			Path:        links.ADMIN_BLOG_POST_MANAGER,
			HTMLHandler: NewBlogPostManagerController().Handler,
		},
		&router.Route{
			Name:        "Admin > Blog > Post Update",
			Path:        links.ADMIN_BLOG_POST_UPDATE,
			HTMLHandler: NewPostUpdateController().Handler,
		},
		&router.Route{
			Name:        "Admin > Blog",
			Path:        links.ADMIN_BLOG,
			HTMLHandler: handler,
		},
		&router.Route{
			Name:        "Admin > Blog > Catchall",
			Path:        links.ADMIN_BLOG + links.CATCHALL,
			HTMLHandler: handler,
		},
	}
}
