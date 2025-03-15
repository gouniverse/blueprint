package widgets

import (
	"net/http"
	"project/app/links"
	"project/config"
	"strings"

	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

var _ Widget = (*blogPostListWidget)(nil) // verify it extends the interface

// == CONSTUCTOR ==============================================================

// NewBlogPostListWidget creates a new instance of the blog post list widget
//
// Parameters:
//   - None
//
// Returns:
//   - *blogPostListWidget - A pointer to the show widget
func NewBlogPostListWidget() *blogPostListWidget {
	return &blogPostListWidget{}
}

// == WIDGET ================================================================

// blogPostListWidget is the struct that renders the blog post list.
//
// This shortcode is used to show the result of the provided content
// if a condition is met.
//
// Example:
// <x-blog-post-list>content</x-blog-post-list>
type blogPostListWidget struct{}
type blogPostListWidgetData struct {
	postList  []blogstore.Post
	page      int
	perPage   int
	postCount int64
}

// == PUBLIC METHODS =========================================================

// Alias the shortcode alias to be used in the template.
func (w *blogPostListWidget) Alias() string {
	return "x-blog-post-list"
}

// Description a user-friendly description of the shortcode.
func (w *blogPostListWidget) Description() string {
	return "Renders a list of the blog posts"
}

// Render implements the shortcode interface.
func (widget *blogPostListWidget) Render(r *http.Request, content string, params map[string]string) string {
	data, errorMessage := widget.prepareData(r)

	if errorMessage != "" {
		return errorMessage
	}

	return widget.postTiles(data).ToHTML()
}

// == PRIVATE METHODS ========================================================
func (widget *blogPostListWidget) postTiles(data blogPostListWidgetData) *hb.Tag {

	url := links.NewWebsiteLinks().Blog(map[string]string{
		"page": "",
	})

	postCountInt, _ := utils.StrToInt(utils.ToString(data.postCount))

	pagination := bs.Pagination(bs.PaginationOptions{
		NumberItems:       postCountInt,
		CurrentPageNumber: data.page,
		PagesToShow:       10,
		PerPage:           data.perPage,
		URL:               url,
	})

	columnCards := lo.Map(data.postList, func(post blogstore.Post, index int) hb.TagInterface {
		postImageURL := post.ImageUrlOrDefault()

		publishedAt := lo.Ternary(post.PublishedAt() == "", "", post.PublishedAtCarbon().Format("d M, Y"))

		postURL := links.NewWebsiteLinks().BlogPost(post.ID(), post.Slug())

		postImage := hb.Image(postImageURL).
			Class("card-img-top rounded-3").
			Style("object-fit: cover;").
			Style("max-height: 180px;").
			Style("aspect-ratio: 9/6;").
			Alt("")

		postTitle := hb.Heading5().
			Class("card-title").
			Style("font-size: 16px; margin-bottom: 10px; text-align: left; font-weight: 800;").
			Text(post.Title())

		postPublished := hb.Paragraph().
			Style("font-size: 12px;	color: #6c757d;	margin-bottom: 20px; text-align: right;").
			Text(publishedAt)

		postSummary := hb.Paragraph().
			Class("card-text").
			Text(post.Summary()).
			Style(`text-align: left;`).
			Style(`font-size: 14px;`).
			Style(`font-weight: 400;`).
			Style(`overflow: hidden;`).
			Style(`text-overflow: ellipsis;`).
			Style(`display: -webkit-box;`).
			Style(`-webkit-line-clamp: 2;`).
			Style(`-webkit-box-orient: vertical;`)

		separator := hb.HR().
			Style(`width: 80%`).
			Style(`margin: 0 auto`).
			Style(`border: 0`).
			Style(`height: 2px`).
			Style(`background-image: linear-gradient(to right, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.75), rgba(0, 0, 0, 0))`)

		card := hb.Div().
			Class("card").
			Style("border: none;").
			Child(postImage).
			Child(hb.Div().
				Class("card-body").
				Style(`padding: 20px 10px;`).
				Child(postTitle).
				Child(postSummary)).
			Child(hb.Div().
				Class("card-footer").
				Style(`background: none;border: none;padding: 0px;`).
				Child(postPublished).
				Child(separator))

		link := hb.Hyperlink().
			Href(postURL).
			// Target("_blank").
			Style("text-decoration: none; color: inherit;").
			Style("display: flex; height: 100%;").
			Child(card)

		return hb.Div().
			Class("col-md-3 col-sm-6 d-flex align-items-stretch").
			Child(link)
	})

	return hb.Section().
		Style("background:#fff;padding-top:40px; padding-bottom: 40px;").
		Children([]hb.TagInterface{
			bs.Container().Children([]hb.TagInterface{
				bs.Row().Class(`g-4`).Children(columnCards),
				hb.Div().Class(`d-flex justify-content-center mt-5 pagination-primary-soft rounded mb-0`).HTML(pagination),
			}),
		})

}

func (widget *blogPostListWidget) prepareData(r *http.Request) (data blogPostListWidgetData, errorMessage string) {
	pageStr := strings.TrimSpace(utils.Req(r, "page", ""))
	page, err := utils.StrToInt(pageStr)

	if err != nil {
		page = 0
	}

	if page < 0 {
		page = 0
	}

	perPage := 12 // 3 rows x 4 postss

	options := blogstore.PostQueryOptions{
		Status:    blogstore.POST_STATUS_PUBLISHED,
		SortOrder: "DESC",
		OrderBy:   "published_at",
		Offset:    page * perPage,
		Limit:     perPage,
	}

	postList, errList := config.BlogStore.PostList(options)

	if errList != nil {
		config.LogStore.ErrorWithContext("Error. At blogController.page", errList.Error())
		return data, "Sorry, there was an error loading the posts. Please try again later."
	}

	postCount, errCount := config.BlogStore.PostCount(options)

	if errCount != nil {
		config.LogStore.ErrorWithContext("Error. At blogController.page", errCount.Error())
		return data, "Sorry, there was an error loading the posts count. Please try again later."
	}

	data.page = page
	data.perPage = perPage
	data.postList = postList
	data.postCount = postCount

	return data, ""
}
