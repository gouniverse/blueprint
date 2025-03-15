package admin

import (
	"net/http"
	"project/app/layouts"
	"project/app/links"
	"project/config"
	"project/pkg/webtheme"

	adminCmsStore "github.com/gouniverse/cmsstore/admin"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
)

type cmsNewController struct {
}

var _ router.ControllerInterface = (*cmsNewController)(nil)

func NewCmsNewController() *cmsNewController {
	return &cmsNewController{}
}

func (controller cmsNewController) Handler(w http.ResponseWriter, r *http.Request) {
	admin, err := adminCmsStore.New(adminCmsStore.AdminOptions{
		Store:                  config.CmsStore,
		Logger:                 &config.Logger,
		BlockEditorDefinitions: webtheme.BlockEditorDefinitions(),
		// BlockEditorRenderer: func(blocks []ui.BlockInterface) string {
		// 	return webtheme.New(blocks).ToHtml()
		// },
		FuncLayout: func(pageTitle string, pageContent string, options struct {
			Styles     []string
			StyleURLs  []string
			Scripts    []string
			ScriptURLs []string
		}) string {
			return layouts.NewAdminLayout(r, layouts.Options{
				Title:      pageTitle + " | CMS (NEW)",
				Content:    hb.Raw(pageContent),
				ScriptURLs: options.ScriptURLs,
				StyleURLs:  options.StyleURLs,
				Scripts:    options.Scripts,
				Styles:     options.Styles,
			}).ToHTML()
		},
		AdminHomeURL: links.NewAdminLinks().Home(map[string]string{}),
	})

	if err != nil {
		config.Logger.Error("At admin > cmsNewController > Handler", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	admin.Handle(w, r)
}
