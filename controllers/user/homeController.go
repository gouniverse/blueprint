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

func (controller *homeController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewUserLayout(r, layouts.Options{
		Request:    r,
		Title:      "Home " + " | User",
		Content:    hb.NewWrap().HTML("You are in user dashboard"),
		StyleURLs:  []string{},
		ScriptURLs: []string{},
		Scripts:    []string{},
		Styles: []string{
			`nav#Toolbar {border-bottom: 4px solid red;}`,
		},
	}).
		ToHTML()
}
