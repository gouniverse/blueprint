package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/gouniverse/router"
)

func routes() []router.Route {
	routes := []router.Route{}
	routes = append(routes, adminRoutes()...)
	routes = append(routes, authRoutes()...)
	routes = append(routes, userRoutes()...)
	routes = append(routes, sharedRoutes()...)
	routes = append(routes, websiteRoutes()...)

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
