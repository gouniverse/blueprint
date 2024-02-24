package website

import (
	"net/http"
	"project/config"
	"strings"

	"github.com/gouniverse/responses"
	"github.com/samber/lo"
)

const CMS_ENABLE_CACHE = false

func NewCmsController() *cmsController {
	return &cmsController{}
}

type cmsController struct{}

func (controller cmsController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	uri := r.RequestURI

	if strings.HasSuffix(uri, ".ico") {
		return ""
	}

	w.Header().Set("Content-Type", "text/html")
	responses.HTMLResponseF(w, r, controller.cmsFrontend)
	return ""
}

// cmsFrontend shows a page from the CMS based on a defined URI
func (controller cmsController) cmsFrontend(w http.ResponseWriter, r *http.Request) string {
	return controller.pageBuldHtmlByAlias(r, r.URL.Path)
}

// pageBuldHTMLByAlias builds the HTML of a page based on alias
func (controller cmsController) pageBuldHtmlByAlias(r *http.Request, alias string) string {
	page, _ := config.Cms.PageFindByAlias(alias)

	if page == nil {
		return "Page with alias '" + alias + "' not found"
	}

	pageAttrs, err := page.GetAttributes()

	if err != nil {
		return "Page '" + alias + "' io exception"
	}

	pageContent := ""
	pageTitle := ""
	pageMetaKeywords := ""
	pageMetaDescription := ""
	pageMetaRobots := ""
	pageCanonicalURL := ""
	pageTemplateID := ""
	for _, attr := range pageAttrs {
		if attr.AttributeKey() == "content" {
			pageContent = attr.AttributeValue()
		}
		if attr.AttributeKey() == "title" {
			pageTitle = attr.AttributeValue()
		}
		if attr.AttributeKey() == "meta_keywords" {
			pageMetaKeywords = attr.AttributeValue()
		}
		if attr.AttributeKey() == "meta_description" {
			pageMetaDescription = attr.AttributeValue()
		}
		if attr.AttributeKey() == "meta_robots" {
			pageMetaRobots = attr.AttributeValue()
		}
		if attr.AttributeKey() == "canonical_url" {
			pageCanonicalURL = attr.AttributeValue()
		}
		if attr.AttributeKey() == "template_id" {
			pageTemplateID = attr.AttributeValue()
		}
	}

	finalContent := lo.If(pageTemplateID == "", pageContent).ElseF(func() string {
		content, _ := config.Cms.TemplateContentFindByID(pageTemplateID)
		return content
	})

	replacements := map[string]string{
		"PageContent":         pageContent,
		"PageCanonicalUrl":    pageCanonicalURL,
		"PageMetaDescription": pageMetaDescription,
		"PageMetaKeywords":    pageMetaKeywords,
		"PageRobots":          pageMetaRobots,
		"PageTitle":           pageTitle,
	}

	for key, value := range replacements {
		finalContent = strings.ReplaceAll(finalContent, "[["+key+"]]", value)
		finalContent = strings.ReplaceAll(finalContent, "[[ "+key+" ]]", value)
	}

	finalContent, _ = config.Cms.ContentRenderBlocks(finalContent)
	finalContent, _ = config.Cms.ContentRenderShortcodes(r, finalContent)
	finalContent, _ = config.Cms.ContentRenderTranslations(finalContent, "en")

	return finalContent
}
