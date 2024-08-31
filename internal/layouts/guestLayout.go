package layouts

import (
	"project/config"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
)

type guestLayout struct {
	title      string
	content    hb.TagInterface
	scriptURLs []string
	scripts    []string
	styleURLs  []string
	styles     []string
}

func NewGuestLayout(options Options) *guestLayout {
	layout := &guestLayout{}
	layout.title = options.Title + " | " + config.AppName
	layout.content = options.Content
	layout.scriptURLs = options.ScriptURLs
	layout.scripts = options.Scripts
	layout.styleURLs = options.StyleURLs
	layout.styles = options.Styles
	return layout
}

func (layout *guestLayout) ToHTML() string {
	layout.styleURLs = append([]string{cdn.BootstrapCss_5_3_3()}, layout.styleURLs...)
	webpage := hb.NewWebpage().
		SetTitle(layout.title).
		SetFavicon(FaviconURL()).
		AddStyles(layout.styles).
		AddStyleURLs(layout.styleURLs).
		AddScripts(layout.scripts).
		AddScriptURLs(layout.scriptURLs).
		AddChild(layout.content)
	return webpage.ToHTML()
}
