package user

import (
	"net/http"
	"project/internal/layouts"

	"github.com/gouniverse/hb"
)

type homeController struct{}

func NewHomeController() *homeController {
	return &homeController{}
}

func (controller *homeController) Handle(w http.ResponseWriter, r *http.Request) string {
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
