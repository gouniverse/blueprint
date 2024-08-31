package website

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"

	"github.com/go-chi/chi/v5"
	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

type blogPostController struct {
}

func NewBlogPostController() *blogPostController {
	return &blogPostController{}
}

func (c blogPostController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
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

	if post.IsUnpublished() {
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

	return layouts.NewGuestLayout(layouts.Options{
		Request:        r,
		WebsiteSection: "Blog.",
		Title:          post.Title(),
		Content:        hb.NewWrap().HTML(c.page(w, r, *post)),
	}).ToHTML()
}

func (c blogPostController) page(w http.ResponseWriter, r *http.Request, post blogstore.Post) string {
	sectionBanner := NewBlogController().sectionBanner()
	return hb.NewWrap().Children([]hb.TagInterface{
		// hb.NewStyle(c.cssSectionIntro()),
		hb.NewStyle(c.css()),
		hb.NewHTML(sectionBanner),
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

func (c *blogPostController) processContent(content string, editor string) string {
	if editor == "BlockArea" {
		return helpers.BlogPostBlocksToString(content)
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
						hb.NewHTML(c.processContent(post.Content(), post.Editor())),
					}),
				}),
				// bs.Column(8).Children([]hb.TagInterface{
				// 	hb.NewDiv().Class("BlogContent").Children([]hb.TagInterface{
				// 		hb.NewHTML(post.Content()),
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

// func (c blogPostController) sectionIntro() *hb.Tag {
// 	sectionIntro := hb.NewSection().ID("SectionIntro").Children([]hb.TagInterface{
// 		bs.Container().Children([]hb.TagInterface{
// 			hb.NewHeading1().HTML("Blog"),
// 		}),
// 	})
// 	return sectionIntro
// }

// func (c blogPostController) cssSectionIntro() string {
// 	imgUrl := utils.PicsumURL(800, 500, utils.PicsumURLOptions{
// 		ID:   82,
// 		Blur: 0,
// 		// Grayscale: true,
// 	})

// 	return `
// 	#SectionIntro {
// 		padding: 80px 0px;
// 		position: relative;
// 		background-image: url(` + imgUrl + `);
// 		background-position: center;
// 		background-size: cover;
// 	}

// 	#SectionIntro h1 {
// 		color: #fff;
// 		font-size: 55px;
// 		font-weight: 700;
// 		letter-spacing: 2px;
// 		margin: 0px 0px 10px 0px;
// 		text-shadow: 2px 2px 2px #333;
// 		text-transform: uppercase;
// 		text-align: center;
// 	}
// 	`
// }
