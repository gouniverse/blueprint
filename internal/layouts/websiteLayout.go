package layouts

import (
	"net/http"

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
	request        *http.Request
	websiteSection string // i.e. Blog, Website
	title          string
	content        *hb.Tag
	scriptURLs     []string
	scripts        []string
	styleURLs      []string
	styles         []string
}

// func (layout *websiteLayout) ToHTML() string {
// layout.styleURLs = append([]string{cdn.BootstrapCss_5_3_0()}, layout.styleURLs...)
// webpage := hb.NewWebpage().
// 	SetTitle(layout.title).
// 	AddStyleURLs(layout.styleURLs).
// 	AddStyles(layout.styles).
// 	AddChild(layout.content)
// return webpage.ToHTML()
// }

// NewWebsiteLayout creates a new WebsiteLayout object.
//
// Parameters:
// - none
//
// Returns:
// - *WebsiteLayout: a pointer to the WebsiteLayout object.
func NewWebsiteOldLayout() *websiteLayout {
	return &websiteLayout{}
}

// type WebsiteLayoutOptions struct {
// 	PageTitle           string
// 	PageContent         string
// 	PageMetaDescription string
// 	PageMetaKeywords    string
// 	PageMetaRobots      string
// 	PageCanonicalUrl    string
// }

// type websiteLayout struct {
// }

func (l *websiteLayout) fnLayout(r *http.Request, content string) string {
	return content
}

// func (l *websiteLayout) Blank(r *http.Request, data WebsiteLayoutOptions) string {
// 	return hb.NewWebpage().
// 		SetTitle(data.PageTitle + " | Website | ProvedExpert.co.uk").
// 		HTML(data.PageContent).
// 		ToHTML()
// }

// func (layout *websiteLayout) ToHTML() string {
// 	mainTemplate, errTemplate := models.NewCmsRepository().TemplateFindByID(config.CmsUserTemplateID)

// 	if errTemplate != nil {
// 		config.Cms.LogStore.ErrorWithContext("At WebsiteLayout", errTemplate.Error())
// 		return "Template error. Please try again later"
// 	}

// 	if mainTemplate == nil {
// 		config.Cms.LogStore.ErrorWithContext("At WebsiteLayout", "template not found")
// 		return "Template not found. Please try again later"
// 	}

// 	if mainTemplate.Status() != models.CMSTEMPLATE_STATUS_ACTIVE {
// 		config.Cms.LogStore.ErrorWithContext("At WebsiteLayout", "template not active")
// 		return "Template not active. Please try again later"
// 	}
// 	content := mainTemplate.Content()

// 	pageTitle := layout.title

// 	if layout.websiteSection != "" {
// 		pageTitle = layout.websiteSection + " | " + layout.title
// 	}

// 	replacements := map[string]string{
// 		"PageContent": layout.fnLayout(layout.request, layout.content.ToHTML()),
// 		// "PageCanonicalUrl":    layout.PageCanonicalUrl,
// 		// "PageMetaDescription": layout.PageMetaDescription,
// 		// "PageMetaKeywords":    layout.PageMetaKeywords,
// 		// "PageRobots":          layout.PageMetaRobots,
// 		"PageTitle": pageTitle,
// 	}

// 	for key, value := range replacements {
// 		content = strings.ReplaceAll(content, "[["+key+"]]", value)
// 		content = strings.ReplaceAll(content, "[[ "+key+" ]]", value)
// 	}

// 	view, err := config.Cms.ContentRenderBlocks(content)
// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At WebsiteLayout", err.Error())
// 	}

// 	view, err = config.Cms.ContentRenderShortcodes(layout.request, view)
// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At WebsiteLayout", err.Error())
// 	}

// 	view, err = config.Cms.ContentRenderTranslations(view, "en")
// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At WebsiteLayout", err.Error())
// 	}

// 	return view
// }
