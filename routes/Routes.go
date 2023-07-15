package routes

import (
	"time"

	"project/config"
	"project/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/gouniverse/responses"
)

// Routes returns the routes of the application
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "text/html", "text/css"))
	router.Use(middleware.GetHead)
	router.Use(middleware.CleanPath)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 30))
	router.Use(httprate.LimitByIP(20, 1*time.Second))  // 20 req per second
	router.Use(httprate.LimitByIP(180, 1*time.Minute)) // 180 req per minute
	router.Use(httprate.LimitByIP(12000, 1*time.Hour)) // 12000 req hour

	// router.Use(middleware.LoggingMiddleware)
	if config.AppEnvironment != config.APP_ENVIRONMENT_TESTING {
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
	}

	if config.AppEnvironment != "testing" {
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
	}

	// Index Controller > Index Page
	router.Handle("/", responses.HTMLHandler(controllers.NewIndexController().AnyIndex))

	// Not Found? Point to the not found controller
	router.Handle("/*", responses.HTMLHandler(controllers.NewPageNotFoundControllerController().AnyIndex))

	return router
}
