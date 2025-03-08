package blog

import (
	"project/internal/links"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {
	blogRoutes := []router.RouteInterface{
		&router.Route{
			Name:        "Guest > Articles",
			Path:        "/articles",
			HTMLHandler: NewBlogController().Handler,
		},
		&router.Route{
			Name:        "Guest > Articles > Post with ID > Index",
			Path:        "/article/{id:[0-9]+}",
			HTMLHandler: NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Articles > Post with ID && Title > Index",
			Path:        "/article/{id:[0-9]+}/{title}",
			HTMLHandler: NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog",
			Path:        links.BLOG,
			HTMLHandler: NewBlogController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog > Post with ID > Index",
			Path:        links.BLOG_POST_WITH_REGEX,
			HTMLHandler: NewBlogPostController().Handler,
		},
		&router.Route{
			Name:        "Guest > Blog > Post with ID && Title > Index",
			Path:        links.BLOG_POST_WITH_REGEX2,
			HTMLHandler: NewBlogPostController().Handler,
		},
	}

	return blogRoutes
}
