package website

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

type blogController struct{}

type blogControllerData struct {
	postList  []blogstore.Post
	postCount int64
	page      int
	perPage   int
}

func NewBlogController() *blogController {
	return &blogController{}
}

func (controller *blogController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareData(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewWebsiteLinks().Home(), 10)
	}

	return layouts.NewGuestLayout(layouts.Options{
		Request:        r,
		WebsiteSection: "Blog",
		Title:          "Recent Posts",
		Content:        hb.NewWrap().HTML(controller.page(data)),
	}).ToHTML()
}

func (controller *blogController) page(data blogControllerData) string {
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

		postImage := hb.NewImage().
			Src(postImageURL).
			Class("card-img-top rounded-3").
			Style("object-fit: cover;").
			Style("max-height: 180px;").
			Style("aspect-ratio: 9/6;").
			Style("border-radius: 0.5rem").
			Alt("")

		postTitle := hb.NewHeading5().
			Class("card-title").
			Style("font-size: 16px; color: #224b8e; margin-bottom: 10px; text-align: left; font-weight: 800;").
			Text(post.Title())

		postPublished := hb.NewParagraph().
			Style("font-size: 12px;	color: #6c757d;	margin-bottom: 20px; text-align: right;").
			Text(publishedAt)

		// postPublished := hb.NewSpan().
		// 	Class(`small`).
		// 	Style(`font-size:12px;color:#666;display:inline-block;padding-right:10px;padding-top:10px;`).
		// 	HTML(publishedAt)

		// postImage := hb.NewDiv().Class(`overflow-hidden rounded-3`).Children([]hb.TagInterface{
		// 	hb.NewImage().
		// 		Class(`card-img`).
		// 		Style(`object-fit:cover;max-height:180px;`).
		// 		Src(postImageURL).
		// 		Alt("course image").
		// 		Attr("loading", "lazy"),
		// 	hb.NewDiv().
		// 		Class(`bg-overlay bg-dark opacity-4`),
		// 	// Badge
		// 	// bs.CardImageTop().Class(`d-flex align-items-start`).Children([]hb.TagInterface{
		// 	// 	hb.NewDiv().Class(`badge text-bg-danger`).Style(`position:absolute;top:10px;left:10px;`).HTML("Student life"),
		// 	// }),
		// })

		postSummary := hb.NewParagraph().
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

		separator := hb.NewHR().
			Style(`width: 80%`).
			Style(`margin: 0 auto`).
			Style(`border: 0`).
			Style(`height: 2px`).
			Style(`background-image: linear-gradient(to right, rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.75), rgba(0, 0, 0, 0))`).
			Style(`opacity: 0.25`).
			Style(`margin-bottom: 20px`)

		card := hb.NewDiv().
			Class("card").
			Style("border: none; margin-bottom: 20px;").
			Child(postImage).
			Child(hb.NewDiv().
				Class("card-body").
				Style(`padding: 20px 10px;`).
				Child(postTitle).
				Child(postSummary)).
			Child(hb.NewDiv().
				Class("card-footer").
				Style(`background: none;border: none;padding: 0px;`).
				Child(postPublished).
				Child(separator))

		link := hb.NewHyperlink().
			Href(postURL).
			Target("_blank").
			Style("text-decoration: none; color: inherit;").
			Style("display: flex; height: 100%;").
			Child(card)

		return hb.NewDiv().
			Class("col-md-3 col-sm-6 d-flex align-items-stretch").
			Child(link)
	})

	section := hb.NewSection().
		Style("background:#fff;padding-top:40px; padding-bottom: 40px;").
		Children([]hb.TagInterface{
			bs.Container().Children([]hb.TagInterface{
				bs.Row().Class(`g-4`).Children(columnCards),
				hb.NewDiv().Class(`d-flex justify-content-center mt-5 pagination-primary-soft rounded mb-0`).HTML(pagination),
			}),
		})

	return hb.NewWrap().Children([]hb.TagInterface{
		hb.NewHTML(controller.sectionBanner()),
		section,
	}).ToHTML()
}

func (controller blogController) prepareData(r *http.Request) (data blogControllerData, errorMessage string) {
	perPage := 12 // 3 rows x 4 postss
	pageStr := strings.TrimSpace(utils.Req(r, "page", ""))
	page, err := utils.StrToInt(pageStr)

	if err != nil {
		page = 0
	}

	if page < 0 {
		page = 0
	}

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

func (controller blogController) sectionBanner() string {
	decorationCross := `
<figure class="position-absolute top-0 start-0" style="width:50px;">	
	<svg width="22px" height="22px" viewBox="0 0 22 22">
		<polygon class="fill-purple" points="22,8.3 13.7,8.3 13.7,0 8.3,0 8.3,8.3 0,8.3 0,13.7 8.3,13.7 8.3,22 13.7,22 13.7,13.7 22,13.7 "></polygon>
	</svg>
</figure>
	`

	decorationStar := `
<figure class="position-absolute top-100 start-50 translate-middle mt-3 ms-n9 d-none d-lg-block">
	<svg>
		<path class="fill-success" d="m181.6 6.7c-0.1 0-0.2-0.1-0.3 0-2.5-0.3-4.9-1-7.3-1.4-2.7-0.4-5.5-0.7-8.2-0.8-1.4-0.1-2.8-0.1-4.1-0.1-0.5 0-0.9-0.1-1.4-0.2-0.9-0.3-1.9-0.1-2.8-0.1-5.4 0.2-10.8 0.6-16.1 1.4-2.7 0.3-5.3 0.8-7.9 1.3-0.6 0.1-1.1 0.3-1.8 0.3-0.4 0-0.7-0.1-1.1-0.1-1.5 0-3 0.7-4.3 1.2-3 1-6 2.4-8.8 3.9-2.1 1.1-4 2.4-5.9 3.9-1 0.7-1.8 1.5-2.7 2.2-0.5 0.4-1.1 0.5-1.5 0.9s-0.7 0.8-1.1 1.2c-1 1-1.9 2-2.9 2.9-0.4 0.3-0.8 0.5-1.2 0.5-1.3-0.1-2.7-0.4-3.9-0.6-0.7-0.1-1.2 0-1.8 0-3.1 0-6.4-0.1-9.5 0.4-1.7 0.3-3.4 0.5-5.1 0.7-5.3 0.7-10.7 1.4-15.8 3.1-4.6 1.6-8.9 3.8-13.1 6.3-2.1 1.2-4.2 2.5-6.2 3.9-0.9 0.6-1.7 0.9-2.6 1.2s-1.7 1-2.5 1.6c-1.5 1.1-3 2.1-4.6 3.2-1.2 0.9-2.7 1.7-3.9 2.7-1 0.8-2.2 1.5-3.2 2.2-1.1 0.7-2.2 1.5-3.3 2.3-0.8 0.5-1.7 0.9-2.5 1.5-0.9 0.8-1.9 1.5-2.9 2.2 0.1-0.6 0.3-1.2 0.4-1.9 0.3-1.7 0.2-3.6 0-5.3-0.1-0.9-0.3-1.7-0.8-2.4s-1.5-1.1-2.3-0.8c-0.2 0-0.3 0.1-0.4 0.3s-0.1 0.4-0.1 0.6c0.3 3.6 0.2 7.2-0.7 10.7-0.5 2.2-1.5 4.5-2.7 6.4-0.6 0.9-1.4 1.7-2 2.6s-1.5 1.6-2.3 2.3c-0.2 0.2-0.5 0.4-0.6 0.7s0 0.7 0.1 1.1c0.2 0.8 0.6 1.6 1.3 1.8 0.5 0.1 0.9-0.1 1.3-0.3 0.9-0.4 1.8-0.8 2.7-1.2 0.4-0.2 0.7-0.3 1.1-0.6 1.8-1 3.8-1.7 5.8-2.3 4.3-1.1 9-1.1 13.3 0.1 0.2 0.1 0.4 0.1 0.6 0.1 0.7-0.1 0.9-1 0.6-1.6-0.4-0.6-1-0.9-1.7-1.2-2.5-1.1-4.9-2.1-7.5-2.7-0.6-0.2-1.3-0.3-2-0.4-0.3-0.1-0.5 0-0.8-0.1s-0.9 0-1.1-0.1-0.3 0-0.3-0.2c0-0.4 0.7-0.7 1-0.8 0.5-0.3 1-0.7 1.5-1l5.4-3.6c0.4-0.2 0.6-0.6 1-0.9 1.2-0.9 2.8-1.3 4-2.2 0.4-0.3 0.9-0.6 1.3-0.9l2.7-1.8c1-0.6 2.2-1.2 3.2-1.8 0.9-0.5 1.9-0.8 2.7-1.6 0.9-0.8 2.2-1.4 3.2-2 1.2-0.7 2.3-1.4 3.5-2.1 4.1-2.5 8.2-4.9 12.7-6.6 5.2-1.9 10.6-3.4 16.2-4 5.4-0.6 10.8-0.3 16.2-0.5h0.5c1.4-0.1 2.3-0.1 1.7 1.7-1.4 4.5 1.3 7.5 4.3 10 3.4 2.9 7 5.7 11.3 7.1 4.8 1.6 9.6 3.8 14.9 2.7 3-0.6 6.5-4 6.8-6.4 0.2-1.7 0.1-3.3-0.3-4.9-0.4-1.4-1-3-2.2-3.9-0.9-0.6-1.6-1.6-2.4-2.4-0.9-0.8-1.9-1.7-2.9-2.3-2.1-1.4-4.2-2.6-6.5-3.5-3.2-1.3-6.6-2.2-10-3-0.8-0.2-1.6-0.4-2.5-0.5-0.2 0-1.3-0.1-1.3-0.3-0.1-0.2 0.3-0.4 0.5-0.6 0.9-0.8 1.8-1.5 2.7-2.2 1.9-1.4 3.8-2.8 5.8-3.9 2.1-1.2 4.3-2.3 6.6-3.2 1.2-0.4 2.3-0.8 3.6-1 0.6-0.2 1.2-0.2 1.8-0.4 0.4-0.1 0.7-0.3 1.1-0.5 1.2-0.5 2.7-0.5 3.9-0.8 1.3-0.2 2.7-0.4 4.1-0.7 2.7-0.4 5.5-0.8 8.2-1.1 3.3-0.4 6.7-0.7 10-1 7.7-0.6 15.3-0.3 23 1.3 4.2 0.9 8.3 1.9 12.3 3.6 1.2 0.5 2.3 1.1 3.5 1.5 0.7 0.2 1.3 0.7 1.8 1.1 0.7 0.6 1.5 1.1 2.3 1.7 0.2 0.2 0.6 0.3 0.8 0.2 0.1-0.1 0.1-0.2 0.2-0.4 0.1-0.9-0.2-1.7-0.7-2.4-0.4-0.6-1-1.4-1.6-1.9-0.8-0.7-2-1.1-2.9-1.6-1-0.5-2-0.9-3.1-1.3-2.5-1.1-5.2-2-7.8-2.8-1-0.8-2.4-1.2-3.7-1.4zm-64.4 25.8c4.7 1.3 10.3 3.3 14.6 7.9 0.9 1 2.4 1.8 1.8 3.5-0.6 1.6-2.2 1.5-3.6 1.7-4.9 0.8-9.4-1.2-13.6-2.9-4.5-1.7-8.8-4.3-11.9-8.3-0.5-0.6-1-1.4-1.1-2.2 0-0.3 0-0.6-0.1-0.9s-0.2-0.6 0.1-0.9c0.2-0.2 0.5-0.2 0.8-0.2 2.3-0.1 4.7 0 7.1 0.4 0.9 0.1 1.6 0.6 2.5 0.8 1.1 0.4 2.3 0.8 3.4 1.1z"></path>
	</svg>
</figure>
	`

	decorationArrow := `
<figure class="position-absolute top-50 end-0 translate-middle-y">
	<svg width="27px" height="27px">
		<path class="fill-orange" d="M13.122,5.946 L17.679,-0.001 L17.404,7.528 L24.661,5.946 L19.683,11.533 L26.244,15.056 L18.891,16.089 L21.686,23.068 L15.400,19.062 L13.122,26.232 L10.843,19.062 L4.557,23.068 L7.352,16.089 L-0.000,15.056 L6.561,11.533 L1.582,5.946 L8.839,7.528 L8.565,-0.001 L13.122,5.946 Z"></path>
	</svg>
</figure>
	`

	homeLink := hb.NewHyperlink().HTML("Home").Href(links.NewWebsiteLinks().Home()).ToHTML()
	blogLink := hb.NewHyperlink().HTML("Blog").Href(links.NewWebsiteLinks().Blog(map[string]string{})).ToHTML()

	return `
<style>
.fill-success {
	fill: #0cbc87 !important;
}
.fill-orange {
	fill: #fd7e14 !important;
}
.fill-purple {
	fill: #6f42c1 !important;
}
</style>
<section class="py-5" style="background:midnightblue;background-image:url(https://sfs.ams3.digitaloceanspaces.com/lesichkov/cms/1378053-bg.jpg)">
	<div class="container">
		<div class="row position-relative">
		    ` + decorationCross + `
			
		
			<!-- Title and breadcrumb -->
			<div class="col-lg-10 mx-auto text-center position-relative">
				` + decorationArrow + `
				` + decorationStar + `
				<!-- Title -->
				<h1 style="color:white;">Blog</h1>
				<!-- Breadcrumb -->
				<div class="d-flex justify-content-center position-relative">
					<nav aria-label="breadcrumb">
						<ol class="breadcrumb mb-0">
							<li class="breadcrumb-item">` + homeLink + `</li>
							<li class="breadcrumb-item active" aria-current="page">` + blogLink + `</li>
						</ol>
					</nav>
				</div>
			</div>
		</div>
	</div>
</section>
	`
}
