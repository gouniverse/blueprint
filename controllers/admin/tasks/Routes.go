package adminTasks

// import (
// 	"net/http"

// 	"project/internal/routes"

// 	queuemanager "project/controllers/admin/tasks/queuemanager"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/gouniverse/router"
// )

// type Route struct {
// 	Path    string
// 	Methods []string
// 	Handler func(w http.ResponseWriter, r *http.Request) string
// }

// func Routes() []router.Route {
// 	routes := []router.Route{
// 		{
// 			Path:        "/",
// 			Methods:     []string{"all"},
// 			HTMLHandler: queuemanager.NewQueueManagerController().AnyIndex,
// 		},
// 		{
// 			Path:        "/*",
// 			Methods:     []string{"all"},
// 			HTMLHandler: queuemanager.NewQueueManagerController().AnyIndex,
// 		},
// 	}

// 	return routes
// }

// func Handler(chiRouter chi.Router) {
// 	routes.ChiRouterHandleRoutes(chiRouter, Routes())
// }
