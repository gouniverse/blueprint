package admin

import (
	"net/http"
	"project/config"
	"project/internal/layouts"

	"github.com/gouniverse/hb"
)

type cmsController struct {
}

func NewCmsController() *cmsController {
	return &cmsController{}
}

func (controller cmsController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	config.Cms.SetFuncLayout(func(content string) string {
		return layouts.NewAdminLayout(r, layouts.Options{
			Title:   "CMS",
			Content: hb.Raw(content),
		}).ToHTML()
	})

	config.Cms.Router(w, r)

	return ""
}
