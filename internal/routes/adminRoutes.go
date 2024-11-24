package routes

import (
	"project/controllers/admin"
	adminBlog "project/controllers/admin/blog"
	adminCms "project/controllers/admin/cms"
	adminFiles "project/controllers/admin/files"
	adminMedia "project/controllers/admin/media"
	adminShop "project/controllers/admin/shop"
	adminStats "project/controllers/admin/stats"
	adminTasks "project/controllers/admin/tasks"
	adminUsers "project/controllers/admin/users"
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
			HTMLHandler: adminCms.NewCmsController().Handler,
		},
		&router.Route{
			Name:    "Admin > Cms New Manager",
			Path:    links.ADMIN_CMS_NEW,
			Handler: adminCms.NewCmsNewController().Handler,
		},
		&router.Route{
			Name:        "Admin > File Manager",
			Path:        links.ADMIN_FILE_MANAGER,
			HTMLHandler: adminFiles.NewFileManagerController().AnyIndex,
		},
		&router.Route{
			Name:        "Admin > Media Manager",
			Path:        links.ADMIN_MEDIA,
			HTMLHandler: adminMedia.NewMediaManagerController().AnyIndex,
		},
	}

	adminRoutes = append(adminRoutes, adminShop.ShopRoutes()...)
	adminRoutes = append(adminRoutes, adminStats.StatsRoutes()...)
	adminRoutes = append(adminRoutes, adminTasks.TaskRoutes()...)
	adminRoutes = append(adminRoutes, adminUsers.UserRoutes()...)

	adminRoutes = append(adminRoutes, []router.RouteInterface{
		&router.Route{
			Name:        "Admin > Home",
			Path:        links.ADMIN_HOME,
			HTMLHandler: admin.NewHomeController().Handler,
		},
		&router.Route{
			Name:        "Admin > Catch All",
			Path:        links.ADMIN_HOME + links.CATCHALL,
			HTMLHandler: admin.NewHomeController().Handler,
		},
	}...)

	router.RoutesPrependMiddlewares(adminRoutes, []router.Middleware{
		middlewares.NewAdminMiddleware(),
	})

	return adminRoutes
}
