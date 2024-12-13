package widgets

import (
	"net/http"
	"project/config"
	"project/internal/helpers"

	// "project/internal/layouts"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
	"github.com/mingrammer/cfmt"
)

var _ Widget = (*blogPostWidget)(nil) // verify it extends the interface

// == CONSTUCTOR ==============================================================

// NewBlogPostWidget creates a new instance of the show widget
//
// Parameters:
//   - None
//
// Returns:
//   - *blogPostWidget - A pointer to the show widget
func NewBlogPostWidget() *blogPostWidget {
	return &blogPostWidget{}
}

// == WIDGET ================================================================

// blogPostWidget is the struct that renders the blog post.
//
// This shortcode is used to show the result of the provided content
// if a condition is met.
//
// Example:
// <x-blog-post>content</x-blog-post>
type blogPostWidget struct{}

// == PUBLIC METHODS =========================================================

// Alias the shortcode alias to be used in the template.
func (w *blogPostWidget) Alias() string {
	return "x-blog-post"
}

// Description a user-friendly description of the shortcode.
func (w *blogPostWidget) Description() string {
	return "Renders a blog post"
}

// Render implements the shortcode interface.
func (w *blogPostWidget) Render(r *http.Request, content string, params map[string]string) string {
	blogsUrl := links.NewWebsiteLinks().Blog(map[string]string{})

	uriParts := strings.Split(r.RequestURI, "/")
	if len(uriParts) < 5 {
		config.LogStore.ErrorWithContext("At blogPostWidget: URI mismatch", map[string]any{"uri": r.RequestURI})
		url := helpers.ToFlashWarningURL("The post you are looking for no longer exists. Redirecting to the blog location...", blogsUrl, 5)
		return hb.Script(`window.location.href = "` + url + `"`).ToHTML()
	}

	postID := uriParts[3]
	postSlug := uriParts[4]

	cfmt.Infoln("BlogPost: ", postID, postSlug)

	if postID == "" {
		config.LogStore.ErrorWithContext("anyPost: post ID is missing", map[string]any{"uri": r.RequestURI})
		url := helpers.ToFlashWarningURL("The post you are looking for no longer exists. Redirecting to the blog location...", blogsUrl, 5)
		return hb.Script(`window.location.href = "` + url + `"`).ToHTML()

	}

	post, errPost := config.BlogStore.PostFindByID(postID)

	if errPost != nil {
		config.LogStore.ErrorWithContext("Error. At BlogPostController.AnyIndex. Post not found", errPost.Error())
		url := helpers.ToFlashWarningURL("The post you are looking for no longer exists. Redirecting to the blog location...", blogsUrl, 5)
		return hb.Script(`window.location.href = "` + url + `"`).ToHTML()
	}

	if post == nil {
		config.LogStore.ErrorWithContext("ERROR: anyPost: post with ID "+postID+" is missing", map[string]any{"postID": postID})
		url := helpers.ToFlashWarningURL("The post you are looking for no longer exists. Redirecting to the blog location...", blogsUrl, 5)
		return hb.Script(`window.location.href = "` + url + `"`).ToHTML()
	}

	if post.IsUnpublished() {
		config.LogStore.WarnWithContext("WARNING: anyPost: post with ID "+postID+" is unpublished", map[string]any{"postID": postID})
		url := helpers.ToFlashWarningURL("The post you are looking for is no longer active. Redirecting to the blog location...", blogsUrl, 5)
		return hb.Script(`window.location.href = "` + url + `"`).ToHTML()
	}

	if postSlug == "" || postSlug != utils.StrSlugify(post.Title(), '-') {
		blogPostURL := links.NewWebsiteLinks().BlogPost(post.ID(), post.Title())
		config.LogStore.ErrorWithContext("ERROR: anyPost: post Title is missing for ID "+postID, "Redirecting to correct URL")
		url := helpers.ToFlashWarningURL("The post location has changed. Redirecting to the new address...", blogPostURL, 5)
		return hb.Script(`window.location.href = "` + url + `"`).ToHTML()
	}

	return w.page(r, *post)
}

// == PRIVATE METHODS ========================================================

func (w *blogPostWidget) page(_ *http.Request, post blogstore.Post) string {
	return hb.Wrap().
		Children([]hb.TagInterface{
			hb.Style(w.css()),
			w.sectionBreadcrumbs(post),
			// NewBlogController().sectionHeader(),
			w.sectionPost(post),
		}).
		ToHTML()
}

func (w *blogPostWidget) css() string {
	return `
#SectionNewsItem {
	padding:50px 0px 80px 0px;
}

#SectionNewsItem .BlogTitle {
	padding:10px 0px 20px 0px;
	font-size:35px;
	font-family: 'Crimson Text', serif;
	color:#224b8e;
	text-align:centre;
}

#SectionNewsItem .BlogContent {
	padding:10px 0px 20px 0px;
	text-align:left;
	font-family:"Times New Roman", Times, serif;
	font-size:20px;
}
	`
}

func (w *blogPostWidget) processContent(content string, editor string) string {
	if editor == "BlockArea" {
		return helpers.BlogPostBlocksToString(content)
	}
	return content
}

func (w *blogPostWidget) sectionBreadcrumbs(post blogstore.Post) *hb.Tag {
	// breadcrumbs := []bs.Breadcrumb{
	// 	{
	// 		Name: "Blog",
	// 		URL:  links.NewWebsiteLinks().Blog(map[string]string{}),
	// 	},
	// 	{
	// 		Name: "Post",
	// 		URL:  links.NewWebsiteLinks().BlogPost(post.ID(), post.Slug()),
	// 	},
	// }

	return hb.Wrap()

	// breadcrumbSection := layouts.NewWebsiteBreadcrumbsSectionWithContainer(breadcrumbs).
	// 	Style("padding: 20px 0;")

	// return breadcrumbSection
}

func (w *blogPostWidget) sectionPost(post blogstore.Post) *hb.Tag {
	sectionPost := hb.Section().
		ID("SectionNewsItem").
		Style(`background:#fff;`).
		Children([]hb.TagInterface{
			bs.Container().Children([]hb.TagInterface{
				bs.Row().Children([]hb.TagInterface{
					bs.Column(12).Children([]hb.TagInterface{
						hb.Div().Class("BlogTitle").Children([]hb.TagInterface{
							hb.Heading1().Style("color:#794FC6;").HTML(post.Title()),
						}),
					}),
				}),
				bs.Row().Children([]hb.TagInterface{
					bs.Column(12).Children([]hb.TagInterface{
						hb.Div().
							Class("BlogImage float-end d-sm-block d-md-inline float-end pt-md-3 pt-lg-3 pb-md-3 pb-lg-3 ps-md-3 ps-lg-3").
							// Style("padding-top:30px; padding-left:30px; padding-bottom:30px;").
							Children([]hb.TagInterface{
								hb.Image(post.ImageUrlOrDefault()).
									Class("img img-responsive img-thumbnail").
									Style("max-width:500px;"),
							}),
						hb.Div().Class("BlogContent").Children([]hb.TagInterface{
							hb.Raw(w.processContent(post.Content(), post.Editor())),
						}),
					}),
				}),
				bs.Row().Style("margin-top:40px;").Children([]hb.TagInterface{
					bs.Column(12).Children([]hb.TagInterface{
						hb.Div().Children([]hb.TagInterface{
							hb.Hyperlink().Class("btn text-white text-center").Style(`background:#1ba1b6;color:#fff;width:600px;max-width:100%;`).Children([]hb.TagInterface{
								// icons.Icon("bi-arrow-left", 16, 16, "#333").Style("margin-right:5px;"),
								hb.Span().HTML("View All Posts"),
							}).Attr("href", links.NewWebsiteLinks().Blog(map[string]string{})),
						}),
					}),
				}),
			}),
		})
	return sectionPost
}
