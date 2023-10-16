package routes

import (
	"net/http"
	"project/config"
	"project/internal/middlewares"

	adminControllers "project/controllers/admin"
	authControllers "project/controllers/auth"
	sharedControllers "project/controllers/shared"
	userControllers "project/controllers/user"
	websiteControllers "project/controllers/website"

	"github.com/go-chi/chi/v5"
	"github.com/gouniverse/dashboard"
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
		{
			Name:    "Admin > Cms Manager",
			Path:    "/admin/cms",
			Handler: adminControllers.NewCmsController().AnyIndex,
		},
		{
			Name:    "Admin > Home Controller > Index Page",
			Path:    "/admin",
			Handler: adminControllers.NewHomeController().AnyIndex,
		},
	}

	authControllers := []router.Route{
		{
			Name:    "Auth > Auth Controller > Index Page",
			Path:    "/auth/auth",
			Handler: authControllers.NewAuthenticationController().AnyIndex,
		},
		{
			Name:    "Auth > Login Controller > Index Page",
			Path:    "/auth/login",
			Handler: authControllers.NewLoginController().AnyIndex,
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

	websiteRoutes := []router.Route{
		{
			Name:    "Website > Cms > Home Page",
			Path:    "/",
			Handler: websiteControllers.NewCmsController().AnyIndex,
		},
		{
			Name:    "Website > Cms > Catch All Pages",
			Path:    "/*",
			Handler: websiteControllers.NewCmsController().AnyIndex,
		},
	}

	sharedRoutes := []router.Route{
		{
			Name:    "Shared > Flash Controller > Index Page",
			Path:    "/flash",
			Handler: sharedControllers.NewFlashController().AnyIndex,
		},
		{
			Name: "Shared > Theme",
			Path: "/theme",
			Handler: func(w http.ResponseWriter, r *http.Request) string {
				dashboard.ThemeHandler(w, r)
				return ""
			},
		},
	}

	routes := []router.Route{}
	routes = append(routes, adminRoutes...)
	routes = append(routes, authControllers...)
	routes = append(routes, userRoutes...)
	routes = append(routes, websiteRoutes...)
	routes = append(routes, sharedRoutes...)

	return routes
}

func RoutesList() (globalMiddlewareList []router.Middleware, routeList []router.Route) {
	return globalMiddlewares(), routes()
}

// Routes returns the routes of the application
func Routes() *chi.Mux {
	globalMiddlewares, routes := RoutesList()
	return router.NewChiRouter(globalMiddlewares, routes)
}
