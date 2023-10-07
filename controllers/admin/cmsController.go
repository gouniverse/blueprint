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
	// config.Cms.SetFuncLayout(func(content string) string {
	// 	return partials.AdminDashboard(r, partials.AdminDashboardOptions{
	// 		Title:   "Home",
	// 		Content: content,
	// 	}).ToHTML()
	// })
	config.Cms.Router(w, r)
	return ""
}
