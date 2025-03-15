package admin

import (
	"project/app/links"

	"github.com/gouniverse/router"
)

func Routes() []router.RouteInterface {

	postCreate := &router.Route{
		Name:        "Admin > Blog > Post Create",
		Path:        links.ADMIN_BLOG_POST_CREATE,
		HTMLHandler: NewPostCreateController().Handler,
	}

	postDelete := &router.Route{
		Name:        "Admin > Blog > Post Delete",
		Path:        links.ADMIN_BLOG_POST_DELETE,
		HTMLHandler: NewPostDeleteController().Handler,
	}

	postManager := &router.Route{
		Name:        "Admin > Blog > Post Manager",
		Path:        links.ADMIN_BLOG_POST_MANAGER,
		HTMLHandler: NewBlogPostManagerController().Handler,
	}

	postUpdate := &router.Route{
		Name:        "Admin > Blog > Post Update",
		Path:        links.ADMIN_BLOG_POST_UPDATE,
		HTMLHandler: NewPostUpdateController().Handler,
	}

	blogHome := &router.Route{
		Name:        "Admin > Blog",
		Path:        links.ADMIN_BLOG,
		HTMLHandler: NewBlogPostManagerController().Handler,
	}

	blogCatchAll := &router.Route{
		Name:        "Admin > Blog > Catch All",
		Path:        links.ADMIN_BLOG + links.CATCHALL,
		HTMLHandler: NewBlogPostManagerController().Handler,
	}

	return []router.RouteInterface{
		postCreate,
		postDelete,
		postManager,
		postUpdate,
		blogHome,
		blogCatchAll,
	}
}
