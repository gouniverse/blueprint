package routes

import (
	"project/controllers/website"
paypalControllers "project/controllers/website/paypal"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"

	blogControllers "project/controllers/website/blog"
	cmsControllers "project/controllers/website/cms"
)

func websiteRoutes() []router.RouteInterface {
	seoRoutes := []router.RouteInterface{
		&router.Route{
			Name: "Website > ads.txt",
			Path: "/ads.txt",
			HTMLHandler: responses.HTMLHandler(func(w http.ResponseWriter, r *http.Request) string {
				//return "google.com, pub-8821108004642146, DIRECT, f08c47fec0942fa0"
				return "google.com, pub-YOURNUMBER, DIRECT, YOURSTRING"
			}),
		},
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

	// !!! Comment these if you use the global routes, as they clash
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
	
	paymentRoutes := []router.RouteInterface{
		&router.Route{
		 	Name:        "Website > Payment Canceled Controller > Handle",
		 	Path:        links.PAYMENT_CANCELED,
		 	HTMLHandler: website.NewPaymentCanceledController().Handle,
		},
		&router.Route{
		 	Name:        "Website > Payment Success Controller > Handle",
		 	Path:        links.PAYMENT_SUCCESS,
		 	HTMLHandler: website.NewPaymentSuccessController().Handle,
		},
		&router.Route{
			Name:        "Guest > Paypal Success Controller > Index",
			Path:        links.PAYPAL_SUCCESS,
			HTMLHandler: paypalControllers.NewPaypalSuccessController().AnyIndex,
		},
		&router.Route{
			Name:        "Guest > Paypal Cancel Controller > Index",
			Path:        links.PAYPAL_CANCEL,
			HTMLHandler: paypalControllers.NewPaypalCancelController().AnyIndex,
		},
		&router.Route{
			Name:        "Guest > Paypal Notify Controller > Index",
			Path:        links.PAYPAL_NOTIFY,
			HTMLHandler: paypalControllers.NewPaypalNotifyController().AnyIndex,
		},
	}

	// !!! Comment these if you use the CMS routes, as they clash
	globalRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Website > Home Controller",
			Path:        links.HOME,
			HTMLHandler: website.HomeController().Handler,
		},
		&router.Route{
			Name:        "Shared > Page Not Found Controller",
			Path:        links.CATCHALL,
			HTMLHandler: shared.PageNotFoundController().Handler,
		},
	}

	websiteRoutes := []router.RouteInterface{}

	// Comment if you do not use the blog routes
	websiteRoutes = append(websiteRoutes, blogRoutes...)

	// Comment if you do not use the payment routes
	// websiteRoutes = append(websiteRoutes, paymentRoutes...)
	websiteRoutes = append(websiteRoutes, seoRoutes...)

	// Comment if you do not use the CMS routes, but global routes
	websiteRoutes = append(websiteRoutes, cmsRoutes...)

	// Comment if you do not use the global routes, but CMS routes
	websiteRoutes = append(websiteRoutes, globalRoutes...)

	return websiteRoutes
}
