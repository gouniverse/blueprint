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
func adminRoutes() []router.Route {
	adminRoutes := []router.Route{
		{
			Name: "Admin > Blog",
			Path: links.ADMIN_BLOG,
			// Handler: adminBlog.NewBlogController().AnyIndex,
			Handler: adminBlog.NewBlogPostManagerController().Handler,
		},
		{
			Name:    "Admin > Blog > Post Create",
			Path:    links.ADMIN_BLOG_POST_CREATE,
			Handler: adminBlog.NewPostCreateController().Handler,
		},
		{
			Name:    "Admin > Blog > Post Delete",
			Path:    links.ADMIN_BLOG_POST_DELETE,
			Handler: adminBlog.NewPostDeleteController().Handler,
		},
		// {
		// 	Name:    "Admin > Blog > Post Details",
		// 	Path:    links.ADMIN_BLOG_POST_VIEW,
		// 	Handler: adminBlog.NewPostViewController().Handler,
		// },
		{
			Name:    "Admin > Blog > Post Manager",
			Path:    links.ADMIN_BLOG_POST_MANAGER,
			Handler: adminBlog.NewBlogPostManagerController().Handler,
		},
		{
			Name:    "Admin > Blog > Post Update",
			Path:    links.ADMIN_BLOG_POST_UPDATE,
			Handler: adminBlog.NewPostUpdateController().Handler,
		},

		{
			Name:    "Admin > Cms Manager",
			Path:    links.ADMIN_CMS,
			Handler: admin.NewCmsController().AnyIndex,
		},
		{
			Name:    "Admin > File Manager",
			Path:    links.ADMIN_MEDIA,
			Handler: adminFiles.NewFileManagerController().AnyIndex,
		},
		{
			Name:    "Admin > Home Controller",
			Path:    links.ADMIN_HOME,
			Handler: admin.NewHomeController().Handler,
		},
	}

	router.RoutesPrependMiddlewares(adminRoutes, []router.Middleware{
		middlewares.NewAdminMiddleware(),
	})

	return adminRoutes
}
