package admin

import (
	"net/http"
	"project/app/layouts"
	"project/config"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
)

type cmsController struct {
}

var _ router.HTMLControllerInterface = (*cmsController)(nil)

func NewCmsController() *cmsController {
	return &cmsController{}
}

func (controller cmsController) Handler(w http.ResponseWriter, r *http.Request) string {
	config.Cms.SetFuncLayout(func(content string) string {
		return layouts.NewAdminLayout(r, layouts.Options{
			Title:   "CMS",
			Content: hb.Raw(content),
		}).ToHTML()
	})

	config.Cms.Router(w, r)

	return ""
}
