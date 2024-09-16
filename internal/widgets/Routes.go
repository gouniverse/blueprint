package widgets

import (
	"github.com/gouniverse/router"
)

const PATH_COMMENTABLE = "/widgets/commentable"

func Routes() []router.RouteInterface {
	return []router.RouteInterface{
		// &router.Route{
		// 	Path:    "/widgets/commentable",
		// 	Handler: NewCommentableWidget().Handler,
		// },
	}
}
