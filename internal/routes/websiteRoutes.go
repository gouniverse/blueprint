package routes

import (
	"project/controllers/website"
	"project/internal/links"

	"github.com/gouniverse/router"

	websiteBlogControllers "project/controllers/website/blog"
)

func websiteRoutes() []router.Route {
	websiteRoutes := []router.Route{
		{
			Name:    "Guest > Articles",
			Path:    "/articles",
			Handler: websiteBlogControllers.NewBlogController().Handler,
		},
		{
			Name:    "Guest > Articles > Post with ID > Index",
			Path:    "/article/{id:[0-9]+}",
			Handler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		{
			Name:    "Guest > Articles > Post with ID && Title > Index",
			Path:    "/article/{id:[0-9]+}/{title}",
			Handler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		{
			Name:    "Guest > Blog",
			Path:    links.BLOG,
			Handler: websiteBlogControllers.NewBlogController().Handler,
		},
		{
			Name:    "Guest > Blog > Post with ID > Index",
			Path:    links.BLOG_POST_WITH_REGEX,
			Handler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		{
			Name:    "Guest > Blog > Post with ID && Title > Index",
			Path:    links.BLOG_POST_WITH_REGEX2,
			Handler: websiteBlogControllers.NewBlogPostController().Handler,
		},
		// {
		// 	Path:    links.HOME,
		// 	Handler: websiteControllers.NewHomeController().AnyIndex,
		// },
		// {
		// 	Path:    links.CATCHALL,
		// 	Handler: sharedControllers.NewPageNotFoundControllerController().AnyIndex,
		// },
		{
			Name:    "Website > Widget Controller > Handler",
			Path:    links.WIDGET,
			Handler: website.NewWidgetController().Handler,
		},
		{
			Name:    "Website > Cms > Home Page",
			Path:    links.HOME,
			Handler: website.NewCmsController().Handler,
		},
		{
			Name:    "Website > Cms > Catch All Pages",
			Path:    links.CATCHALL,
			Handler: website.NewCmsController().Handler,
		},
	}

	return websiteRoutes
}
