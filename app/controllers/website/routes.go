package website

import (
	"net/http"
	"project/config"
	"project/internal/links"

	"github.com/gouniverse/responses"
	"github.com/gouniverse/router"

	"project/app/controllers/shared"
	blogControllers "project/app/controllers/website/blog"

	cms "project/app/controllers/website/cms"
	seo "project/app/controllers/website/seo"
	// paypalControllers "project/controllers/website/paypal"
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
		HTMLHandler: newHomeController().Handler,
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

	// paymentSuccess := &router.Route{
	// 	Name:        "Website > Payment Success Controller",
	// 	Path:        links.PAYMENT_SUCCESS,
	// 	HTMLHandler: payment.NewPaymentSuccessController().Handler,
	// }

	// paymentCancel := &router.Route{
	// 	Name:        "Website > Payment Cancel Controller",
	// 	Path:        links.PAYMENT_CANCELED,
	// 	HTMLHandler: payment.NewPaymentCanceledController().Handler,
	// }

	// These are custom routes for the website, that cannot be served by the CMS
	websiteRoutes := []router.RouteInterface{
		faviconRoute,
		contactRoute,
		contactSubmitRoute,
		// paymentSuccess,
		// paymentCancel,
	}

	// Comment if you do not use the blog routes
	websiteRoutes = append(websiteRoutes, blogRoutes()...)

	// Comment if you do not use the payment routes
	// websiteRoutes = append(websiteRoutes, paymentRoutes...)
	websiteRoutes = append(websiteRoutes, seoRoutes()...)

	if config.CmsStoreUsed {
		websiteRoutes = append(websiteRoutes, cms.Routes()...)
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
	adsRoute := &router.Route{
		Name: "Website > ads.txt",
		Path: "/ads.txt",
		HTMLHandler: responses.HTMLHandler(func(w http.ResponseWriter, r *http.Request) string {
			//return "google.com, pub-8821108004642146, DIRECT, f08c47fec0942fa0"
			return "google.com, pub-YOURNUMBER, DIRECT, YOURSTRING"
		}),
	}

	robotsRoute := &router.Route{
		Name:        "Website > RobotsTxt",
		Path:        "/robots.txt",
		HTMLHandler: seo.NewRobotsTxtController().Handler,
	}

	securityRoute := &router.Route{
		Name:        "Website > SecurityTxt",
		Path:        "/security.txt",
		HTMLHandler: seo.NewSecurityTxtController().Handler,
	}

	sitemapRoute := &router.Route{
		Name:        "Website > Sitemap",
		Path:        "/sitemap.xml",
		HTMLHandler: seo.NewSitemapXmlController().Handler,
	}

	return []router.RouteInterface{
		adsRoute,
		robotsRoute,
		securityRoute,
		sitemapRoute,
	}
}
