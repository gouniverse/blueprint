package layouts

import (
	"net/http"
	"project/config"

	"github.com/gouniverse/hb"
)

func NewWebsiteLayout(options Options) *websiteLayout {
	layout := &websiteLayout{}
	layout.title = options.Title // + " | " + config.AppName
	layout.content = options.Content
	layout.scriptURLs = options.ScriptURLs
	layout.scripts = options.Scripts
	layout.styleURLs = options.StyleURLs
	layout.styles = options.Styles
	return layout
}

type websiteLayout struct {
	request *http.Request
	// websiteSection string // i.e. Blog, Website
	title      string
	content    hb.TagInterface
	scriptURLs []string
	scripts    []string
	styleURLs  []string
	styles     []string
}

func (layout *websiteLayout) ToHTML() string {
	html, err := config.Cms.TemplateRenderHtmlByID(layout.request, config.CmsUserTemplateID, struct {
		PageContent         string
		PageCanonicalURL    string
		PageMetaDescription string
		PageMetaKeywords    string
		PageMetaRobots      string
		PageTitle           string
		Language            string
	}{
		PageContent:         layout.content.ToHTML(),
		PageCanonicalURL:    "",
		PageMetaDescription: "",
		PageMetaKeywords:    "",
		PageMetaRobots:      "",
		PageTitle:           layout.title,
		Language:            "en",
	})

	if err != nil {
		config.Logger.Error("At WebsiteLayout", "error", err.Error())
		return "Template error. Please try again later"
	}

	return html
}
