package layouts

import (
	"net/http"
	"project/config"
	"project/pkg/cmsstore"
	"strings"

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
	content        hb.TagInterface
	scriptURLs     []string
	scripts        []string
	styleURLs      []string
	styles         []string
}

// NewWebsiteLayout creates a new WebsiteLayout object.
//
// Parameters:
// - none
//
// Returns:
// - *WebsiteLayout: a pointer to the WebsiteLayout object.
// func NewWebsiteOldLayout() *websiteLayout {
// 	return &websiteLayout{}
// }

func (l *websiteLayout) fnLayout(_ *http.Request, content string) string {
	return content
}

func (layout *websiteLayout) ToHTML() string {
	mainTemplate, errTemplate := cmsstore.NewCmsRepository().TemplateFindByID(config.CmsUserTemplateID)

	if errTemplate != nil {
		config.LogStore.ErrorWithContext("At WebsiteLayout", errTemplate.Error())
		return "Template error. Please try again later"
	}

	if mainTemplate == nil {
		config.LogStore.ErrorWithContext("At WebsiteLayout", "template not found")
		return "Template not found. Please try again later"
	}

	if mainTemplate.Status() != cmsstore.CMSTEMPLATE_STATUS_ACTIVE {
		config.LogStore.ErrorWithContext("At WebsiteLayout", "template not active")
		return "Template not active. Please try again later"
	}
	content := mainTemplate.Content()

	pageTitle := layout.title

	if layout.websiteSection != "" {
		pageTitle = layout.websiteSection + " | " + layout.title
	}

	replacements := map[string]string{
		"PageContent": layout.fnLayout(layout.request, layout.content.ToHTML()),
		// "PageCanonicalUrl":    layout.PageCanonicalUrl,
		// "PageMetaDescription": layout.PageMetaDescription,
		// "PageMetaKeywords":    layout.PageMetaKeywords,
		// "PageRobots":          layout.PageMetaRobots,
		"PageTitle": pageTitle,
	}

	for key, value := range replacements {
		content = strings.ReplaceAll(content, "[["+key+"]]", value)
		content = strings.ReplaceAll(content, "[[ "+key+" ]]", value)
	}

	view, err := config.Cms.ContentRenderBlocks(content)
	if err != nil {
		config.LogStore.ErrorWithContext("At WebsiteLayout", err.Error())
	}

	view, err = config.Cms.ContentRenderShortcodes(layout.request, view)
	if err != nil {
		config.LogStore.ErrorWithContext("At WebsiteLayout", err.Error())
	}

	view, err = config.Cms.ContentRenderTranslations(view, "en")
	if err != nil {
		config.LogStore.ErrorWithContext("At WebsiteLayout", err.Error())
	}

	return view
}
