package routes

import (
	"project/config"
	"project/internal/middlewares"

	adminControllers "project/controllers/admin"
	authControllers "project/controllers/auth"
	guestControllers "project/controllers/guest"
	sharedControllers "project/controllers/shared"
	userControllers "project/controllers/user"

	"github.com/go-chi/chi/v5"
	"github.com/gouniverse/router"
)

func globalMiddlewares() []router.Middleware {
	globalMiddlewares := []router.Middleware{
		router.NewCompressMiddleware(5, "text/html", "text/css"),
		router.NewGetHeadMiddleware(),
		router.NewCleanPathMiddleware(),
		router.NewRedirectSlashesMiddleware(),
		router.NewTimeoutMiddleware(30),                 // 30s timeout
		router.NewRateLimitByIpMiddleware(20, 1),        // 20 req per second
		router.NewRateLimitByIpMiddleware(180, 1*60),    // 180 req per minute
		router.NewRateLimitByIpMiddleware(12000, 60*60), // 12000 req hour
	}

	if config.AppEnvironment != config.APP_ENVIRONMENT_TESTING {
		globalMiddlewares = append(globalMiddlewares, router.NewLoggerMiddleware())
		globalMiddlewares = append(globalMiddlewares, router.NewRecovererMiddleware())
	}

	return globalMiddlewares
}

func routes() []router.Route {
	adminRoutes := []router.Route{
		// Enable if CMS is used
		// {
		// 	// Admin > Cms
		// 	Name:    "Admin > Cms",
		// 	Path:    "/admin/cms",
		// 	Handler: adminControllers.NewCmsController().AnyIndex,
		// },
		{
			// Admin > Home Controller > Index Page
			Name:    "Admin > Home Controller > Index Page",
			Path:    "/admin",
			Handler: adminControllers.NewHomeController().AnyIndex,
		},
	}

	authControllers := []router.Route{
		{
			// Auth > Home Controller > Index Page
			Name:    "Auth > Home Controller > Index Page",
			Path:    "/auth",
			Handler: authControllers.NewHomeController().AnyIndex,
		},
	}

	userRoutes := []router.Route{
		{
			// User > Home Controller > Index Page
			Name:    "User > Home Controller > Index Page",
			Path:    "/user",
			Handler: userControllers.NewHomeController().AnyIndex,
		},
	}

	router.RoutesPrependMiddlewares(userRoutes, []router.Middleware{
		middlewares.NewUserMiddleware(),
	})

	router.RoutesPrependMiddlewares(adminRoutes, []router.Middleware{
		middlewares.NewAdminMiddleware(),
	})

	guestRoutes := []router.Route{
		{
			// Guest > Home Controller > Index Page
			Name:    "Guest > Home Controller > Index Page",
			Path:    "/",
			Handler: guestControllers.NewHomeController().AnyIndex,
		},
		// Enable if CMS is used
		// {
		// 	// Guest > Cms
		// 	Name:    "Guest > Cms",
		// 	Path:    "/",
		// 	Handler: guestControllers.NewCmsController().AnyIndex,
		// },
	}

	sharedRoutes := []router.Route{
		{
			// Not Found? Point to the not found controller
			Path:    "/*",
			Handler: sharedControllers.NewPageNotFoundControllerController().AnyIndex,
		},
		// Enable if CMS is used
		// {
		// 	// Guest > Cms
		// 	Name:    "Guest > Cms",
		// 	Path:    "/*",
		// 	Handler: guestControllers.NewCmsController().AnyIndex,
		// },
	}

	routes := []router.Route{}
	routes = append(routes, adminRoutes...)
	routes = append(routes, authControllers...)
	routes = append(routes, guestRoutes...)
	routes = append(routes, userRoutes...)
	routes = append(routes, sharedRoutes...)

	return routes
}

// Routes returns the routes of the application
func Routes() *chi.Mux {
	return router.NewChiRouter(globalMiddlewares(), routes())
}
