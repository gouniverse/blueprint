package layouts

import (
	"project/config"

	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
)

// == TYPE ===================================================================

// GuestLayout is a layout for guest pages, which are not connected to the CMS
// it differs from the website layout, which uses the CMS template
//
// NOTE: It is used for the registration page, which only has a
// registration form and no navigation
type guestLayout struct {
	title      string
	content    hb.TagInterface
	scriptURLs []string
	scripts    []string
	styleURLs  []string
	styles     []string
}

// == CONSTRUCTOR =============================================================

// NewGuestLayout creates a new guest layout
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

// == PUBLIC METHODS ==========================================================

// ToHTML generates the HTML for the guest layout
func (layout *guestLayout) ToHTML() string {
	layout.styleURLs = append([]string{cdn.BootstrapCss_5_3_3()}, layout.styleURLs...)
	webpage := hb.Webpage().
		SetTitle(layout.title).
		SetFavicon(FaviconURL()).
		AddStyles(layout.styles).
		AddStyleURLs(layout.styleURLs).
		AddScripts(layout.scripts).
		AddScriptURLs(layout.scriptURLs).
		AddChild(layout.content)
	return webpage.ToHTML()
}
