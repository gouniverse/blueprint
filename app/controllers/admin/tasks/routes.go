package admin

import (
	"project/app/links"

	"github.com/gouniverse/router"
)

func TaskRoutes() []router.RouteInterface {
	return []router.RouteInterface{
		// &router.Route{
		// 	Name:        "Admin > Users > User Create",
		// 	Path:        links.ADMIN_USERS_USER_CREATE,
		// 	HTMLHandler: NewUserCreateController().Handler,
		// },
		// &router.Route{
		// 	Name:        "Admin > Users > User Delete",
		// 	Path:        links.ADMIN_USERS_USER_DELETE,
		// 	HTMLHandler: NewUserDeleteController().Handler,
		// },
		// &router.Route{
		// 	Name:        "Admin > Users > User Impersonate",
		// 	Path:        links.ADMIN_USERS_USER_IMPERSONATE,
		// 	HTMLHandler: NewUserImpersonateController().Handler,
		// },
		// &router.Route{
		// 	Name:        "Admin > Users > User Manager",
		// 	Path:        links.ADMIN_USERS_USER_MANAGER,
		// 	HTMLHandler: NewUserManagerController().Handler,
		// },
		// &router.Route{
		// 	Name:        "Admin > Users > User Update",
		// 	Path:        links.ADMIN_USERS_USER_UPDATE,
		// 	HTMLHandler: NewUserUpdateController().Handler,
		// },
		&router.Route{
			Name:        "Admin > Tasks > Home",
			Path:        links.ADMIN_TASKS,
			HTMLHandler: TaskController().Handler,
		},
		&router.Route{
			Name:        "Admin > Tasks > Catchall",
			Path:        links.ADMIN_TASKS + links.CATCHALL,
			HTMLHandler: TaskController().Handler,
		},
	}
}
