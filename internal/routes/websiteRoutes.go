package routes

import (
	"project/controllers/website"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"

	websiteBlogControllers "project/controllers/website/blog"
)

func websiteRoutes() []router.RouteInterface {
	websiteRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Guest > Articles",
			Path:        "/articles",
			HTMLHandler: websiteBlogControllers.NewBlogController().Handler,
		},
		&router.Route{
			Name:        "Guest > Articles > Post with ID > Index",
			Path:        "/article/{id:[0-9]+}",
			HTMLHandler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Articles > Post with ID && Title > Index",
			Path:        "/article/{id:[0-9]+}/{title}",
			HTMLHandler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog",
			Path:        links.BLOG,
			HTMLHandler: websiteBlogControllers.NewBlogController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog > Post with ID > Index",
			Path:        links.BLOG_POST_WITH_REGEX,
			HTMLHandler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog > Post with ID && Title > Index",
			Path:        links.BLOG_POST_WITH_REGEX2,
			HTMLHandler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		// {
		// 	Path:    links.HOME,
		// 	Handler: websiteControllers.NewHomeController().AnyIndex,
		// },
		// {
		// 	Path:    links.CATCHALL,
		// 	Handler: sharedControllers.NewPageNotFoundControllerController().AnyIndex,
		// },
		// &router.Route{
		// 	Name:        "Website > Payment Canceled Controller > Handle",
		// 	Path:        links.PAYMENT_CANCELED,
		// 	HTMLHandler: websitePayment.NewPaymentCanceledController().Handle,
		// },
		// &router.Route{
		// 	Name:        "Website > Payment Success Controller > Handle",
		// 	Path:        links.PAYMENT_SUCCESS,
		// 	HTMLHandler: websitePayment.NewPaymentSuccessController().Handle,
		// },
		&router.Route{
			Name:        "Website > RobotsTxt",
			Path:        "/robots.txt",
			HTMLHandler: website.NewRobotsTxtController().Handler,
		},
		&router.Route{
			Name:        "Website > SecurityTxt",
			Path:        "/security.txt",
			HTMLHandler: website.NewSecurityTxtController().Handler,
		},
		&router.Route{
			Name:        "Website > Sitemap",
			Path:        "/sitemap.xml",
			HTMLHandler: website.NewSitemapXmlController().Handler,
		},
		&router.Route{
			Name:        "Website > Widget Controller > Handler",
			Path:        links.WIDGET,
			HTMLHandler: website.NewWidgetController().Handler,
		},
		&router.Route{
			Name:        "Website > Cms > Home Page",
			Middlewares: []router.Middleware{middlewares.NewStatsMiddleware()},
			Path:        links.HOME,
			HTMLHandler: website.NewCmsController().Handler,
		},
		&router.Route{
			Name:        "Website > Cms > Catch All Pages",
			Middlewares: []router.Middleware{middlewares.NewStatsMiddleware()},
			Path:        links.CATCHALL,
			HTMLHandler: website.NewCmsController().Handler,
		},
	}

	return websiteRoutes
}
