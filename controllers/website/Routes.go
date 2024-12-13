package website

import (
	"net/http"
	"project/config"
	"project/internal/links"
	"project/internal/middlewares"

	"github.com/gouniverse/router"

	"project/controllers/shared"
	blogControllers "project/controllers/website/blog"
	cmsControllers "project/controllers/website/cms"
)

func Routes() []router.RouteInterface {
	routes := []router.RouteInterface{}
	routes = append(routes, websiteRoutes()...)
	return routes
}

func websiteRoutes() []router.RouteInterface {
	homeRoute := &router.Route{
		Name:        "Website > Home Controller",
		Path:        links.HOME,
		HTMLHandler: HomeController().Handler,
	}

	pageNotFoundRoute := &router.Route{
		Name:        "Shared > Page Not Found Controller",
		Path:        links.CATCHALL,
		HTMLHandler: shared.PageNotFoundController().Handler,
	}

	faviconRoute := &router.Route{
		Name: "Website Favicon",
		Path: "/favicon.svg",
		HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
			w.Header().Add("Content-Type", "image/svg+xml .svg .svgz")
			return `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 32 32"><circle cx="20" cy="8" r="1" fill="currentColor"></circle><circle cx="23" cy="8" r="1" fill="currentColor"></circle><circle cx="26" cy="8" r="1" fill="currentColor"></circle><path d="M28 4H4a2.002 2.002 0 0 0-2 2v20a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2zm0 2v4H4V6zM4 12h6v14H4zm8 14V12h16v14z" fill="currentColor"></path></svg>`
		},
	}

	contactRoute := &router.Route{
		Path:        links.CONTACT,
		Methods:     []string{http.MethodGet, http.MethodPost},
		HTMLHandler: NewContactController().AnyIndex,
	}

	contactSubmitRoute := &router.Route{
		Path:        links.CONTACT,
		Methods:     []string{http.MethodPost},
		HTMLHandler: NewContactController().AnyIndex,
	}

	// These are custom routes for the website, that cannot be served by the CMS
	websiteRoutes := []router.RouteInterface{
		faviconRoute,
		contactRoute,
		contactSubmitRoute,
	}

	// Comment if you do not use the blog routes
	websiteRoutes = append(websiteRoutes, blogRoutes()...)

	// Comment if you do not use the payment routes
	// websiteRoutes = append(websiteRoutes, paymentRoutes...)

	websiteRoutes = append(websiteRoutes, seoRoutes()...)

	if config.CmsStoreUsed {
		websiteRoutes = append(websiteRoutes, cmsRoutes()...)
	} else {
		websiteRoutes = append(websiteRoutes, homeRoute)
		websiteRoutes = append(websiteRoutes, pageNotFoundRoute)
	}

	return websiteRoutes
}

func blogRoutes() []router.RouteInterface {
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
	}

	return blogRoutes
}

func cmsRoutes() []router.RouteInterface {
	return []router.RouteInterface{
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
}

// func paymentRoutes() []router.RouteInterface {
// 	paymentRoutes := []router.RouteInterface{
// 		&router.Route{
// 			Name:        "Website > Payment Canceled Controller > Handle",
// 			Path:        links.PAYMENT_CANCELED,
// 			HTMLHandler: website.NewPaymentCanceledController().Handle,
// 		},
// 		&router.Route{
// 			Name:        "Website > Payment Success Controller > Handle",
// 			Path:        links.PAYMENT_SUCCESS,
// 			HTMLHandler: website.NewPaymentSuccessController().Handle,
// 		},
// 		&router.Route{
// 			Name:        "Guest > Paypal Success Controller > Index",
// 			Path:        links.PAYPAL_SUCCESS,
// 			HTMLHandler: paypalControllers.NewPaypalSuccessController().AnyIndex,
// 		},
// 		&router.Route{
// 			Name:        "Guest > Paypal Cancel Controller > Index",
// 			Path:        links.PAYPAL_CANCEL,
// 			HTMLHandler: paypalControllers.NewPaypalCancelController().AnyIndex,
// 		},
// 		&router.Route{
// 			Name:        "Guest > Paypal Notify Controller > Index",
// 			Path:        links.PAYPAL_NOTIFY,
// 			HTMLHandler: paypalControllers.NewPaypalNotifyController().AnyIndex,
// 		},
// 	}

// 	return paymentRoutes
// }

func seoRoutes() []router.RouteInterface {
	robotsRoute := &router.Route{
		Name:        "Website > RobotsTxt",
		Path:        "/robots.txt",
		HTMLHandler: NewRobotsTxtController().Handler,
	}

	securityRoute := &router.Route{
		Name:        "Website > SecurityTxt",
		Path:        "/security.txt",
		HTMLHandler: NewSecurityTxtController().Handler,
	}

	sitemapRoute := &router.Route{
		Name:        "Website > Sitemap",
		Path:        "/sitemap.xml",
		HTMLHandler: NewSitemapXmlController().Handler,
	}

	return []router.RouteInterface{
		robotsRoute,
		securityRoute,
		sitemapRoute,
	}
}
