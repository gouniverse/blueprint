package user

import (
	"net/http"
	"project/internal/layouts"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
)

// == CONTROLLER ==============================================================

type homeController struct{}

var _ router.HTMLControllerInterface = (*homeController)(nil)

// == CONSTRUCTOR =============================================================

func NewHomeController() *homeController {
	return &homeController{}
}

// == PUBLIC METHODS ==========================================================

func (controller *homeController) Handler(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewUserLayout(r, layouts.Options{
		Request:    r,
		Title:      "Home",
		Content:    controller.view(),
		StyleURLs:  []string{},
		ScriptURLs: []string{},
		Scripts:    []string{},
		Styles:     []string{},
	}).
		ToHTML()
}

func (controller *homeController) view() *hb.Tag {
	return hb.NewWrap().HTML("You are in user dashboard")
}
