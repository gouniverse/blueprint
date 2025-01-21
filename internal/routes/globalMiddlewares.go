package routes

import (
	"project/config"
	"project/internal/middlewares"

	"github.com/gouniverse/router"
)

func globalMiddlewares() []router.Middleware {
	globalMiddlewares := []router.Middleware{
		middlewares.NewJailBotsMiddleware(),
		router.NewCompressMiddleware(5, "text/html", "text/css"),
		router.NewGetHeadMiddleware(),
		router.NewCleanPathMiddleware(),
		router.NewRedirectSlashesMiddleware(),
		//router.NewNakedDomainToWwwMiddleware([]string{"localhost", "127.0.0.1", "http://sinevia.local"}),
		router.NewTimeoutMiddleware(30),                 // 30s timeout
		router.NewRateLimitByIpMiddleware(20, 1),        // 20 req per second
		router.NewRateLimitByIpMiddleware(180, 1*60),    // 180 req per minute
		router.NewRateLimitByIpMiddleware(12000, 60*60), // 12000 req hour
	}

	if config.AppEnvironment != config.APP_ENVIRONMENT_TESTING {
		globalMiddlewares = append(globalMiddlewares, router.NewLoggerMiddleware())
		globalMiddlewares = append(globalMiddlewares, router.NewRecovererMiddleware())
	}

	globalMiddlewares = append(globalMiddlewares, middlewares.NewLogRequestMiddleware())
	globalMiddlewares = append(globalMiddlewares, middlewares.NewThemeMiddleware())
	globalMiddlewares = append(globalMiddlewares, middlewares.NewAuthMiddleware())

	return globalMiddlewares
}
