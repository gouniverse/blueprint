package website

import (
	"bytes"
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type blogPostController struct {
}

func NewBlogPostController() *blogPostController {
	return &blogPostController{}
}

var _ router.HTMLControllerInterface = (*blogPostController)(nil)

func (c blogPostController) Handler(w http.ResponseWriter, r *http.Request) string {
	postID := chi.URLParam(r, "id")
	postSlug := chi.URLParam(r, "title")
	blogsUrl := links.NewWebsiteLinks().Blog(map[string]string{})

	if postID == "" {
		config.LogStore.ErrorWithContext("anyPost: post ID is missing", map[string]any{"uri": r.RequestURI})
		helpers.ToFlash(w, r, "warning", "The post you are looking for no longer exists. Redirecting to the blog location...", blogsUrl, 5)
		return "post is missing"
	}

	post, errPost := config.BlogStore.PostFindByID(postID)

	if errPost != nil {
		config.LogStore.ErrorWithContext("Error. At BlogPostController.AnyIndex. Post not found", errPost.Error())
		helpers.ToFlash(w, r, "warning", "The post you are looking for no longer exists. Redirecting to the blog location...", blogsUrl, 5)
		return "post is missing"
	}

	if post == nil {
		config.LogStore.ErrorWithContext("ERROR: anyPost: post with ID "+postID+" is missing", map[string]any{"postID": postID})
		helpers.ToFlash(w, r, "warning", "The post you are looking for no longer exists. Redirecting to the blog location...", blogsUrl, 5)
		return ""
	}

	if !c.accessAllowed(r, *post) {
		config.LogStore.WarnWithContext("WARNING: anyPost: post with ID "+postID+" is unpublished", map[string]any{"postID": postID})
		helpers.ToFlash(w, r, "warning", "The post you are looking for is no longer active. Redirecting to the blog location...", blogsUrl, 5)
		return ""
	}

	if postSlug == "" || postSlug != utils.StrSlugify(post.Title(), '-') {
		url := links.NewWebsiteLinks().BlogPost(post.ID(), post.Title())
		config.LogStore.ErrorWithContext("ERROR: anyPost: post Title is missing for ID "+postID, "Redirecting to correct URL")
		helpers.ToFlash(w, r, "success", "The post location has changed. Redirecting to the new address...", url, 5)
		return ""

		// return guestshared.WebsiteTemplate(guestshared.WebsiteTemplateOptions{
		// 	Title:           post.Title() + " | Blog Post",
		// 	MetaDescription: post.MetaDescription(),
		// 	Content:         c.htmlPost(*post) + sharedWidget.HTML(),
		// 	Scripts: []string{
		// 		// uncdn.NotifyJs(),
		// 		c.commentsJs(helpers.SafeDereference(post)),
		// 		sharedWidget.JS(),
		// 	},
		// 	Styles: []string{
		// 		c.cssSectionIntro(),
		// 		c.cssPost(),
		// 		sharedWidget.CSS(),
		// 	},
		// 	HTTPRequest: r,
		// })
	}

	return layouts.NewWebsiteLayout(layouts.Options{
		Request:        r,
		WebsiteSection: "Blog.",
		Title:          post.Title(),
		Content:        hb.NewWrap().HTML(c.page(*post)),
	}).ToHTML()
}

func (controller blogPostController) accessAllowed(r *http.Request, post blogstore.Post) bool {
	if post.IsPublished() {
		return true // everyone can access published posts
	}

	authUser := helpers.GetAuthUser(r)

	// If the user is not logged in, they can't access unpublished posts
	if authUser == nil {
		return false
	}

	// If the user is an administrator, they can access unpublished posts
	if authUser.IsAdministrator() {
		return true
	}

	return false // default to false
}

func (c blogPostController) page(post blogstore.Post) string {
	sectionBanner := NewBlogController().sectionBanner()
	return hb.NewWrap().Children([]hb.TagInterface{
		// hb.NewStyle(c.cssSectionIntro()),
		hb.NewStyle(c.css()),
		hb.Raw(sectionBanner),
		// c.sectionIntro(),
		c.sectionPost(post),
	}).ToHTML()
}

// css returns the CSS code for the blog post section.
//
// No parameters.
// Returns a string containing the CSS code.
func (c blogPostController) css() string {
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

func (controller *blogPostController) processContent(content string, editor string) string {
	if editor == blogstore.POST_EDITOR_BLOCKAREA {
		return helpers.BlogPostBlocksToString(content)
	}
	if editor == blogstore.POST_EDITOR_MARKDOWN {
		return controller.markdownToHtml(content)
	}
	return content
}

func (c *blogPostController) sectionPost(post blogstore.Post) *hb.Tag {
	// nextPost, _ := models.NewBlogRepository().PostFindNext(post)
	// prevPost, _ := models.NewBlogRepository().PostFindPrevious(post)
	sectionPost := hb.NewSection().ID("SectionNewsItem").Style(`background:#fff;`).Children([]hb.TagInterface{
		bs.Container().Children([]hb.TagInterface{
			bs.Row().Children([]hb.TagInterface{
				bs.Column(12).Children([]hb.TagInterface{
					hb.NewDiv().Class("BlogTitle").Children([]hb.TagInterface{
						hb.NewHeading1().Style("color:#794FC6;").HTML(post.Title()),
					}),
				}),
			}),
			bs.Row().Children([]hb.TagInterface{
				bs.Column(12).Children([]hb.TagInterface{
					hb.NewDiv().Class("BlogImage float-end").Style("padding-top:30px; padding-left:30px; padding-bottom:30px; width:600px; max-width:100%;").Children([]hb.TagInterface{
						hb.NewImage().Class("img img-responsive img-thumbnail").Attrs(map[string]string{
							"src": post.ImageUrlOrDefault(),
						}),
					}),
					hb.NewDiv().Class("BlogContent").Children([]hb.TagInterface{
						hb.Raw(c.processContent(post.Content(), post.Editor())),
					}),
				}),
				// bs.Column(8).Children([]hb.TagInterface{
				// 	hb.NewDiv().Class("BlogContent").Children([]hb.TagInterface{
				// 		hb.Raw(post.Content()),
				// 	}),
				// }),
				// bs.Column(4).Children([]hb.TagInterface{
				// 	hb.NewDiv().Class("BlogImage").Children([]hb.TagInterface{
				// 		hb.NewImage().Class("img img-responsive img-thumbnail").Attrs(map[string]string{
				// 			"src": post.ImageUrlOrDefault(),
				// 		}),
				// 	}),
				// }),
			}),
			// bs.Row().Children([]hb.TagInterface{
			// 	bs.Column(6).Children([]hb.TagInterface{
			// 		lo.IfF(prevPost != nil, func() *hb.Tag {
			// 			link := links.NewWebsiteLinks().BlogPost(prevPost.ID(), prevPost.Title(), map[string]string{})
			// 			return hb.NewDiv().Children([]hb.TagInterface{
			// 				hb.NewHyperlink().Children([]hb.TagInterface{
			// 					icons.Icon("bi-chevron-left", 20, 20, "#333").Style("margin-right:5px;"),
			// 					hb.NewSpan().HTML("Previous"),
			// 				}).Attr("href", link).
			// 					Style("font-weight:bold; font-size:20px;"),
			// 				hb.NewDiv().HTML(prevPost.Title()),
			// 			})
			// 		}).ElseF(func() *hb.Tag {
			// 			return hb.NewSpan().HTML("")
			// 		}),
			// 	}),
			// 	bs.Column(6).Children([]hb.TagInterface{
			// 		lo.IfF(nextPost != nil, func() *hb.Tag {
			// 			link := links.NewWebsiteLinks().BlogPost(nextPost.ID(), nextPost.Title(), map[string]string{})
			// 			return hb.NewDiv().Children([]hb.TagInterface{
			// 				hb.NewHyperlink().Children([]hb.TagInterface{
			// 					hb.NewSpan().HTML("Next"),
			// 					icons.Icon("bi-chevron-right", 20, 20, "#333").Style("margin-right:5px;"),
			// 				}).Attr("href", link).
			// 					Style("font-weight:bold; font-size:20px;"),
			// 				hb.NewDiv().HTML(nextPost.Title()),
			// 			}).Style("text-align:right;")
			// 		}).ElseF(func() *hb.Tag {
			// 			return hb.NewSpan().HTML("")
			// 		}),
			// 	}),
			// }),
			bs.Row().Style("margin-top:40px;").Children([]hb.TagInterface{
				bs.Column(12).Children([]hb.TagInterface{
					hb.NewDiv().Children([]hb.TagInterface{
						hb.NewHyperlink().Class("btn text-white text-center").Style(`background:#1ba1b6;color:#fff;width:600px;max-width:100%;`).Children([]hb.TagInterface{
							// icons.Icon("bi-arrow-left", 16, 16, "#333").Style("margin-right:5px;"),
							hb.NewSpan().HTML("View All Posts"),
						}).Attr("href", links.NewWebsiteLinks().Blog(map[string]string{})),
					}),
				}),
			}),
		}),
	})
	return sectionPost
}

// markdownToHtml converts a markdown text to html
//
// 1. the text is trimmed of any white spaces
// 2. if the text is empty, it returns an empty string
// 3. the text is converted to html using the goldmark library
func (controller *blogPostController) markdownToHtml(text string) string {
	text = strings.TrimSpace(text)

	if text == "" {
		return ""
	}

	var buf bytes.Buffer
	md := goldmark.New(
		// goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			// html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert([]byte(text), &buf); err != nil {
		panic(err)
	}

	return buf.String()
}
