package admin

import (
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

	// subscriptionPlans := &router.Route{
	// 	Name:        "Admin > Subscription Plans Controller > Index",
	// 	Path:        links.ADMIN_SUBSCRIPTION_PLANS,
	// 	HTMLHandler: NewSubscriptionPlanController().AnyIndex,
	// }

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
