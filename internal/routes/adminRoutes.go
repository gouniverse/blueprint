package routes

import (
	"project/controllers/admin"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"
)

// adminRoutes these are the routes for the administrator
func adminRoutes() []router.Route {
	adminRoutes := []router.Route{
		{
			Name:    "Admin > Blog",
			Path:    "/admin/blog",
			Handler: admin.NewBlogController().AnyIndex,
		},
		{
			Name:    "Admin > Cms Manager",
			Path:    links.ADMIN_CMS,
			Handler: admin.NewCmsController().AnyIndex,
		},
		{
			Name:    "Admin > File Manager",
			Path:    "/admin/filemanager",
			Handler: admin.NewFileMangerController().AnyIndex,
		},
		{
			Name:    "Admin > Home Controller > Index Page",
			Path:    links.ADMIN_HOME,
			Handler: admin.NewHomeController().AnyIndex,
		},
	}

	router.RoutesPrependMiddlewares(adminRoutes, []router.Middleware{
		middlewares.NewAdminMiddleware(),
	})

	return adminRoutes
}
