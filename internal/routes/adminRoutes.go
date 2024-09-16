package routes

import (
	"project/controllers/admin"
	adminBlog "project/controllers/admin/blog"
	adminFiles "project/controllers/admin/files"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"
)

// adminRoutes these are the routes for the administrator
func adminRoutes() []router.RouteInterface {
	adminRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Admin > Blog",
			Path:        links.ADMIN_BLOG,
			HTMLHandler: adminBlog.NewBlogPostManagerController().Handler,
		},
		&router.Route{
			Name:        "Admin > Blog > Post Create",
			Path:        links.ADMIN_BLOG_POST_CREATE,
			HTMLHandler: adminBlog.NewPostCreateController().Handler,
		},
		&router.Route{
			Name:        "Admin > Blog > Post Delete",
			Path:        links.ADMIN_BLOG_POST_DELETE,
			HTMLHandler: adminBlog.NewPostDeleteController().Handler,
		},
		// {
		// 	Name:    "Admin > Blog > Post Details",
		// 	Path:    links.ADMIN_BLOG_POST_VIEW,
		// 	Handler: adminBlog.NewPostViewController().Handler,
		// },
		&router.Route{
			Name:        "Admin > Blog > Post Manager",
			Path:        links.ADMIN_BLOG_POST_MANAGER,
			HTMLHandler: adminBlog.NewBlogPostManagerController().Handler,
		},
		&router.Route{
			Name:        "Admin > Blog > Post Update",
			Path:        links.ADMIN_BLOG_POST_UPDATE,
			HTMLHandler: adminBlog.NewPostUpdateController().Handler,
		},
		&router.Route{
			Name:        "Admin > Cms Manager",
			Path:        links.ADMIN_CMS,
			HTMLHandler: admin.NewCmsController().AnyIndex,
		},
		&router.Route{
			Name:        "Admin > File Manager",
			Path:        links.ADMIN_MEDIA,
			HTMLHandler: adminFiles.NewFileManagerController().AnyIndex,
		},
		&router.Route{
			Name:        "Admin > Home Controller",
			Path:        links.ADMIN_HOME,
			HTMLHandler: admin.NewHomeController().Handler,
		},
	}

	router.RoutesPrependMiddlewares(adminRoutes, []router.Middleware{
		middlewares.NewAdminMiddleware(),
	})

	return adminRoutes
}
