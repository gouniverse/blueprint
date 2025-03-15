package routes

import (
	"project/app/controllers/admin"
	"project/app/controllers/auth"
	"project/app/controllers/shared"
	"project/app/controllers/user"
	"project/app/controllers/website"
	"project/app/widgets"

	"github.com/go-chi/chi/v5"

	"github.com/gouniverse/router"
)

func routes() []router.RouteInterface {
	routes := []router.RouteInterface{}

	routes = append(routes, admin.Routes()...)
	routes = append(routes, auth.Routes()...)
	routes = append(routes, shared.Routes()...)
	routes = append(routes, user.Routes()...)
	routes = append(routes, widgets.Routes()...)
	routes = append(routes, website.Routes()...)

	return routes
}

func RoutesList() (globalMiddlewareList []router.Middleware, routeList []router.RouteInterface) {
	return globalMiddlewares(), routes()
}

// Routes returns the routes of the application
func Routes() *chi.Mux {
	globalMiddlewares, routes := RoutesList()
	return router.NewChiRouter(globalMiddlewares, routes)
}
