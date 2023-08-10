package routes

import (
	"net/http"
	"project/config"
	"project/controllers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/gouniverse/router"
)

// Routes returns the routes of the application
func Routes() *chi.Mux {
	globalMiddlewares := []func(http.Handler) http.Handler{
		middleware.Compress(5, "text/html", "text/css"),
		middleware.GetHead,
		middleware.CleanPath,
		middleware.RedirectSlashes,
		middleware.Timeout(time.Second * 30),
		httprate.LimitByIP(20, 1*time.Second),  // 20 req per second
		httprate.LimitByIP(180, 1*time.Minute), // 180 req per minute
		httprate.LimitByIP(12000, 1*time.Hour), // 12000 req hour
	}

	if config.AppEnvironment != config.APP_ENVIRONMENT_TESTING {
		globalMiddlewares = append(globalMiddlewares, middleware.Logger)
		globalMiddlewares = append(globalMiddlewares, middleware.Recoverer)
	}

	routes := []router.Route{
		// Index Controller > Index Page
		{
			Path:    "/",
			Handler: controllers.NewIndexController().AnyIndex,
		},
		// Not Found? Point to the not found controller
		{
			Path:    "/*",
			Handler: controllers.NewPageNotFoundControllerController().AnyIndex,
		},
	}

	return router.NewChiRouter(globalMiddlewares, routes)
}
