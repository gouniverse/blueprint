package admin

import (
	adminBlog "project/app/controllers/admin/blog"
	adminCms "project/app/controllers/admin/cms"
	adminFiles "project/app/controllers/admin/files"
	adminMedia "project/app/controllers/admin/media"
	adminShop "project/app/controllers/admin/shop"
	adminStats "project/app/controllers/admin/stats"
	adminTasks "project/app/controllers/admin/tasks"
	adminUsers "project/app/controllers/admin/users"
	"project/app/links"
	"project/app/middlewares"

	"github.com/gouniverse/router"
)

// Routes these are the routes for the administrator
func Routes() []router.RouteInterface {
	home := &router.Route{
		Name:        "Admin > Home",
		Path:        links.ADMIN_HOME,
		HTMLHandler: NewHomeController().Handler,
	}

	homeCatchAll := &router.Route{
		Name:        "Admin > Catch All",
		Path:        links.ADMIN_HOME + links.CATCHALL,
		HTMLHandler: NewHomeController().Handler,
	}

	adminRoutes := []router.RouteInterface{}
	adminRoutes = append(adminRoutes, adminBlog.Routes()...)
	adminRoutes = append(adminRoutes, adminCms.Routes()...)
	adminRoutes = append(adminRoutes, adminFiles.Routes()...)
	adminRoutes = append(adminRoutes, adminMedia.Routes()...)
	adminRoutes = append(adminRoutes, adminShop.ShopRoutes()...)
	adminRoutes = append(adminRoutes, adminStats.StatsRoutes()...)
	adminRoutes = append(adminRoutes, adminTasks.TaskRoutes()...)
	adminRoutes = append(adminRoutes, adminUsers.UserRoutes()...)
	// adminRoutes = append(adminRoutes, []router.RouteInterface{subscriptionPlans}...)
	adminRoutes = append(adminRoutes, []router.RouteInterface{home, homeCatchAll}...)

	router.RoutesPrependMiddlewares(adminRoutes, []router.Middleware{
		middlewares.NewAdminMiddleware(),
	})

	return adminRoutes
}
