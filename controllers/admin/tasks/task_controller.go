package admin

import (
	"log/slog"
	"net/http"
	"project/config"
	"project/internal/layouts"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/samber/lo"

	taskAdmin "github.com/gouniverse/taskstore/admin"
)

func TaskController() router.HTMLControllerInterface {
	return &taskController{
		logger: config.Logger,
	}
}

type taskController struct {
	logger slog.Logger
}

func (c *taskController) Handler(w http.ResponseWriter, r *http.Request) string {
	uptimeAdminUi, err := taskAdmin.UI(taskAdmin.UIOptions{
		ResponseWriter: w,
		Request:        r,
		Logger:         &config.Logger,
		Store:          config.TaskStore,
		Layout:         &adminLayout{},
	})

	ui := lo.IfF(err != nil, func() hb.TagInterface {
		c.logger.Error("At admin > taskController > Handler", "error", err.Error())
		return hb.Raw(err.Error())
	}).ElseF(func() hb.TagInterface {
		return uptimeAdminUi
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

var _ taskAdmin.Layout = (*adminLayout)(nil)
