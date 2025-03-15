package admin

import (
	"log/slog"
	"net/http"
	"project/app/layouts"
	"project/config"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/samber/lo"

	statsAdmin "github.com/gouniverse/statsstore/admin"
)

func StatsController() router.HTMLControllerInterface {
	return &statsController{
		logger: config.Logger,
	}
}

type statsController struct {
	logger slog.Logger
}

func (c *statsController) Handler(w http.ResponseWriter, r *http.Request) string {
	statsUI, err := statsAdmin.UI(statsAdmin.UIOptions{
		ResponseWriter: w,
		Request:        r,
		Logger:         &config.Logger,
		Store:          config.StatsStore,
		Layout:         &adminLayout{},
		HomeURL:        "/admin",
		WebsiteUrl:     "https://lesichkov.co.uk",
	})

	ui := lo.IfF(err != nil, func() hb.TagInterface {
		c.logger.Error("At admin > statsController > Handler", "error", err.Error())
		return hb.Raw(err.Error())
	}).ElseF(func() hb.TagInterface {
		return statsUI
	})

	return ui.ToHTML()
}

type adminLayout struct {
	title string
	body  string

	scriptURLs []string
	scripts    []string

	styleURLs []string
	styles    []string
}

func (a *adminLayout) SetTitle(title string) {
	a.title = title
}

func (a *adminLayout) SetBody(body string) {
	a.body = body
}

func (a *adminLayout) SetScriptURLs(urls []string) {
	a.scriptURLs = urls
}

func (a *adminLayout) SetScripts(scripts []string) {
	a.scripts = scripts
}

func (a *adminLayout) SetStyleURLs(urls []string) {
	a.styleURLs = urls
}

func (a *adminLayout) SetStyles(styles []string) {
	a.styles = styles
}

func (a *adminLayout) Render(w http.ResponseWriter, r *http.Request) string {
	return layouts.NewAdminLayout(r, layouts.Options{
		Title:      a.title,
		Content:    hb.Raw(a.body),
		ScriptURLs: a.scriptURLs,
		Scripts:    a.scripts,
		StyleURLs:  a.styleURLs,
		Styles:     a.styles,
	}).ToHTML()
}

var _ statsAdmin.Layout = (*adminLayout)(nil)
