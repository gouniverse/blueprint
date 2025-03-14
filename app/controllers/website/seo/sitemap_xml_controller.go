package website

import (
	"net/http"
	"project/config"
	"project/internal/links"
	"strings"

	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/responses"
	"github.com/gouniverse/router"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
)

type sitemapXmlController struct{}

// NewSitemapXmlController creates a new instance of the sitemapXmlController struct.
func NewSitemapXmlController() *sitemapXmlController {
	return &sitemapXmlController{}
}

var _ router.HTMLControllerInterface = (*sitemapXmlController)(nil)

func (c sitemapXmlController) Handler(w http.ResponseWriter, r *http.Request) string {
	responses.XMLResponseF(w, r, c.buildSitemapXML)
	return ""
}

func (c sitemapXmlController) buildSitemapXML(w http.ResponseWriter, r *http.Request) string {
	locations := []string{
		"/",
		// "/about",
		// "/contact",
		// "/faq",
		// "/marketplace",
		"/robots.txt",
		// "/privacy-policy",
		// "/terms-of-use",
	}
	postList, err := config.BlogStore.PostList(blogstore.PostQueryOptions{
		Status:    blogstore.POST_STATUS_PUBLISHED,
		OrderBy:   "title",
		SortOrder: sb.DESC,
		Limit:     1000,
	})

	if err != nil {
		config.LogStore.ErrorWithContext("At sitemapXmlController > anySitemapXML", err.Error())
		return ""
	}

	lo.ForEach(postList, func(post blogstore.Post, index int) {
		locations = append(locations, links.NewWebsiteLinks().BlogPost(post.ID(), post.Title()))
	})

	timeNow := carbon.Now().ToIso8601String()

	xml := `<?xml version="1.0" encoding="UTF-8"?>`
	xml += `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
	lo.ForEach(locations, func(location string, index int) {
		url := lo.Ternary(strings.HasPrefix(location, "http"), location, links.RootURL()+location)

		priority := "0.80"
		if index == 0 {
			priority = "1.00"
		}

		xml += "<url>"
		xml += "<loc>" + url + "</loc>"
		xml += "<lastmod>" + timeNow + "</lastmod>"
		xml += "<priority>" + priority + "</priority>"
		xml += "</url>"
	})
	xml += "</urlset>"

	return xml
}
