package routes

import (
	"project/controllers/website"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"

	blogControllers "project/controllers/website/blog"
	cmsControllers "project/controllers/website/cms"
)

func websiteRoutes() []router.RouteInterface {
	blogRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Guest > Articles",
			Path:        "/articles",
			HTMLHandler: blogControllers.NewBlogController().Handler,
		},
		&router.Route{
			Name:        "Guest > Articles > Post with ID > Index",
			Path:        "/article/{id:[0-9]+}",
			HTMLHandler: blogControllers.NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Articles > Post with ID && Title > Index",
			Path:        "/article/{id:[0-9]+}/{title}",
			HTMLHandler: blogControllers.NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog",
			Path:        links.BLOG,
			HTMLHandler: blogControllers.NewBlogController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog > Post with ID > Index",
			Path:        links.BLOG_POST_WITH_REGEX,
			HTMLHandler: blogControllers.NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog > Post with ID && Title > Index",
			Path:        links.BLOG_POST_WITH_REGEX2,
			HTMLHandler: blogControllers.NewBlogPostController().Handler,
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

	}

	seoRoutes := []router.RouteInterface{
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
	}

	cmsRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Website > Widget Controller > Handler",
			Path:        links.WIDGET,
			HTMLHandler: cmsControllers.NewWidgetController().Handler,
		},
		&router.Route{
			Name:        "Website > Cms > Home Page",
			Middlewares: []router.Middleware{middlewares.NewStatsMiddleware()},
			Path:        links.HOME,
			HTMLHandler: cmsControllers.NewCmsController().Handler,
		},
		&router.Route{
			Name:        "Website > Cms > Catch All Pages",
			Middlewares: []router.Middleware{middlewares.NewStatsMiddleware()},
			Path:        links.CATCHALL,
			HTMLHandler: cmsControllers.NewCmsController().Handler,
		},
	}

	websiteRoutes := []router.RouteInterface{}
	websiteRoutes = append(websiteRoutes, blogRoutes...)
	websiteRoutes = append(websiteRoutes, seoRoutes...)
	websiteRoutes = append(websiteRoutes, cmsRoutes...)

	return websiteRoutes
}
