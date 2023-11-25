package admin

import (
	"net/http"
	"project/config"
)

type cmsController struct {
}

func NewCmsController() *cmsController {
	return &cmsController{}
}

func (controller cmsController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	config.Cms.SetFuncLayout(func(content string) string {
		return layout(r, layoutOptions{
			Title:   "CMS",
			Content: content,
		}).ToHTML()
	})
	config.Cms.Router(w, r)
	return ""
}
