package admin

import (
	"net/http"
	"project/config"

	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"
	"strings"

	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

// == CONTROLLER ==============================================================

type blogPostManagerController struct{}

var _ router.HTMLControllerInterface = (*blogPostManagerController)(nil)

// == CONSTRUCTOR =============================================================

func NewBlogPostManagerController() *blogPostManagerController {
	return &blogPostManagerController{}
}

func (controller *blogPostManagerController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareData(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewAdminLinks().Home(), 10)
	}

	return layouts.NewAdminLayout(r, layouts.Options{
		Title:   "Blog | Post Manager",
		Content: controller.page(data),
		ScriptURLs: []string{
			cdn.Htmx_1_9_4(),
			cdn.Sweetalert2_10(),
		},
		Styles: []string{},
	}).ToHTML()
}

func (controller *blogPostManagerController) page(data blogPostManagerControllerData) hb.TagInterface {
	breadcrumbs := layouts.Breadcrumbs([]layouts.Breadcrumb{
		{
			Name: "Home",
			URL:  links.NewAdminLinks().Home(),
		},
		{
			Name: "Blog",
			URL:  links.NewAdminLinks().BlogPostManager(map[string]string{}),
		},
		{
			Name: "Post Manager",
			URL:  links.NewAdminLinks().BlogPostManager(map[string]string{}),
		},
	})

	buttonPostNew := hb.NewButton().
		Class("btn btn-primary float-end").
		Child(hb.NewI().Class("bi bi-plus-circle").Style("margin-top:-4px;margin-right:8px;font-size:16px;")).
		HTML("New Post").
		HxGet(links.NewAdminLinks().BlogPostCreate(map[string]string{})).
		HxTarget("body").
		HxSwap("beforeend")

	title := hb.NewHeading1().
		HTML("Blog. Post Manager").
		Child(buttonPostNew)

	return hb.NewDiv().
		Class("container").
		Child(title).
		Child(breadcrumbs).
		Child(controller.tablePosts(data))
}

func (controller *blogPostManagerController) prepareData(r *http.Request) (data blogPostManagerControllerData, errorMessage string) {
	var err error

	data.page = utils.Req(r, "page", "0")
	data.pageInt = cast.ToInt(data.page)
	data.perPage = cast.ToInt(utils.Req(r, "per_page", "10"))
	data.sortOrder = utils.Req(r, "sort_order", sb.DESC)
	data.sortBy = utils.Req(r, "by", blogstore.COLUMN_CREATED_AT)
	data.status = utils.Req(r, "status", "")
	data.search = utils.Req(r, "search", "")
	data.dateFrom = utils.Req(r, "date_from", carbon.Now().AddYears(-1).ToDateString())
	data.dateTo = utils.Req(r, "date_to", carbon.Now().ToDateString())
	data.customerID = utils.Req(r, "customer_id", "")

	query := blogstore.PostQueryOptions{
		Search:               data.search,
		Offset:               data.pageInt * data.perPage,
		Limit:                data.perPage,
		Status:               data.status,
		CreatedAtGreaterThan: data.dateFrom + " 00:00:00",
		CreatedAtLessThan:    data.dateTo + " 23:59:59",
		SortOrder:            data.sortOrder,
		OrderBy:              data.sortBy,
	}

	data.blogList, err = config.BlogStore.
		// EnableDebug(true).
		PostList(query)

	if err != nil {
		config.LogStore.ErrorWithContext("At blogPostManagerController > prepareData", err.Error())
		return data, "error retrieving posts"
	}

	// DEBUG: cfmt.Successln("Invoice List: ", blogList)

	data.blogCount, err = config.BlogStore.
		// EnableDebug().
		PostCount(query)

	if err != nil {
		config.LogStore.ErrorWithContext("At blogPostManagerController > prepareData", err.Error())
		return data, "Error retrieving posts count"
	}

	return data, ""
}

func (controller *blogPostManagerController) tablePosts(data blogPostManagerControllerData) hb.TagInterface {
	table := hb.NewTable().
		Class("table table-striped table-hover table-bordered").
		Children([]hb.TagInterface{
			hb.NewThead().Children([]hb.TagInterface{
				hb.NewTR().Children([]hb.TagInterface{
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "Post", "title")).
						Text(", ").
						Child(controller.sortableColumnLabel(data, "Reference", "title")).
						Style(`cursor: pointer;`),
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "Status", "status")).
						Style("width: 200px;cursor: pointer;"),
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "Featured", "featured")).
						Style("width: 1px;cursor: pointer;"),
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "Published", "published_at")).
						Style("width: 1px;cursor: pointer;"),
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "Created", "created_at")).
						Style("width: 1px;cursor: pointer;"),
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "Modified", "updated_at")).
						Style("width: 1px;cursor: pointer;"),
					hb.NewTH().
						HTML("Actions"),
				}),
			}),
			hb.NewTbody().Children(lo.Map(data.blogList, func(blog blogstore.Post, _ int) hb.TagInterface {
				blogLink := hb.NewHyperlink().
					HTML(blog.Title()).
					Href(links.NewWebsiteLinks().BlogPost(blog.ID(), blog.Slug())).
					Target("_blank")

				status := hb.NewSpan().
					Style(`font-weight: bold;`).
					StyleIf(blog.IsPublished(), `color:green;`).
					StyleIf(blog.IsTrashed(), `color:silver;`).
					StyleIf(blog.IsDraft(), `color:blue;`).
					StyleIf(blog.IsUnpublished(), `color:red;`).
					HTML(blog.Status())

				buttonEdit := hb.NewHyperlink().
					Class("btn btn-primary me-2").
					Child(hb.NewI().Class("bi bi-pencil-square")).
					Title("Edit").
					Href(links.NewAdminLinks().BlogPostUpdate(map[string]string{"post_id": blog.ID()})).
					Target("_blank")

				buttonDelete := hb.NewHyperlink().
					Class("btn btn-danger").
					Child(hb.NewI().Class("bi bi-trash")).
					Title("Delete").
					HxGet(links.NewAdminLinks().BlogPostDelete(map[string]string{"post_id": blog.ID()})).
					HxTarget("body").
					HxSwap("beforeend")

				return hb.NewTR().Children([]hb.TagInterface{
					hb.NewTD().
						Child(hb.NewDiv().Child(blogLink)).
						Child(hb.NewDiv().
							Style("font-size: 11px;").
							HTML("Ref: ").
							HTML(blog.ID())),
					hb.NewTD().
						Child(status),
					hb.NewTD().
						HTML(blog.Featured()),
					hb.NewTD().
						Child(hb.NewDiv().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(blog.PublishedAtCarbon().Format("d M Y"))),
					hb.NewTD().
						Child(hb.NewDiv().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(blog.CreatedAtCarbon().Format("d M Y"))),
					hb.NewTD().
						Child(hb.NewDiv().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(blog.UpdatedAtCarbon().Format("d M Y"))),
					hb.NewTD().
						Child(buttonEdit).
						Child(buttonDelete),
				})
			})),
		})

	// cfmt.Successln("Table: ", table)

	return hb.NewWrap().Children([]hb.TagInterface{
		controller.tableFilter(data),
		table,
		controller.tablePagination(data, int(data.blogCount), data.pageInt, data.perPage),
	})
}

func (controller *blogPostManagerController) sortableColumnLabel(data blogPostManagerControllerData, tableLabel string, columnName string) hb.TagInterface {
	isSelected := strings.EqualFold(data.sortBy, columnName)

	direction := lo.If(data.sortOrder == "asc", "desc").Else("asc")

	if !isSelected {
		direction = "asc"
	}

	link := links.NewAdminLinks().BlogPostManager(map[string]string{
		"page":        "0",
		"by":          columnName,
		"sort":        direction,
		"date_from":   data.dateFrom,
		"date_to":     data.dateTo,
		"status":      data.status,
		"search":      data.search,
		"customer_id": data.customerID,
	})
	return hb.NewHyperlink().
		HTML(tableLabel).
		Child(controller.sortingIndicator(columnName, data.sortBy, direction)).
		Href(link)
}

func (controller *blogPostManagerController) sortingIndicator(columnName string, sortByColumnName string, sortOrder string) hb.TagInterface {
	isSelected := strings.EqualFold(sortByColumnName, columnName)

	direction := lo.If(isSelected && sortOrder == "asc", "up").
		ElseIf(isSelected && sortOrder == "desc", "down").
		Else("none")

	sortingIndicator := hb.NewSpan().
		Class("sorting").
		HTMLIf(direction == "up", "&#8595;").
		HTMLIf(direction == "down", "&#8593;").
		HTMLIf(direction != "down" && direction != "up", "")

	return sortingIndicator
}

func (controller *blogPostManagerController) tableFilter(data blogPostManagerControllerData) hb.TagInterface {
	statusList := []map[string]string{
		{"id": "", "name": "All Statuses"},
		{"id": blogstore.POST_STATUS_DRAFT, "name": "Draft"},
		{"id": blogstore.POST_STATUS_UNPUBLISHED, "name": "Unpublished"},
		{"id": blogstore.POST_STATUS_PUBLISHED, "name": "Published"},
		{"id": blogstore.POST_STATUS_TRASH, "name": "Deleted"},
	}

	searchButton := hb.NewButton().
		Type("submit").
		Child(hb.NewI().Class("bi bi-search")).
		Class("btn btn-primary w-100 h-100")

	period := hb.NewDiv().Class("form-group").Children([]hb.TagInterface{
		hb.NewLabel().
			HTML("Period").
			Style("margin-bottom: 0px;"),
		hb.NewDiv().Class("input-group").Children([]hb.TagInterface{
			hb.NewInput().
				Type(hb.TYPE_DATE).
				Name("date_from").
				Value(data.dateFrom).
				OnChange("FORM_TRANSACTIONS.submit()").
				Class("form-control"),
			hb.NewSpan().
				HTML(" : ").
				Class("input-group-text"),
			hb.NewInput().
				Type(hb.TYPE_DATE).
				Name("date_to").
				Value(data.dateTo).
				OnChange("FORM_TRANSACTIONS.submit()").
				Class("form-control"),
		}),
	})

	search := hb.NewDiv().Class("form-group").Children([]hb.TagInterface{
		hb.NewLabel().
			HTML("Search").
			Style("margin-bottom: 0px;"),
		hb.NewInput().
			Type("search").
			Name("search").
			Value(data.search).
			Class("form-control").
			Placeholder("reference, title, content ..."),
	})

	status := hb.NewDiv().Class("form-group").Children([]hb.TagInterface{
		hb.NewLabel().
			HTML("Status").
			Style("margin-bottom: 0px;"),
		hb.NewSelect().
			Name("status").
			Class("form-select").
			OnChange("FORM_TRANSACTIONS.submit()").
			Children(lo.Map(statusList, func(status map[string]string, index int) hb.TagInterface {
				return hb.NewOption().
					Value(status["id"]).
					HTML(status["name"]).
					AttrIf(data.status == status["id"], "selected", "selected")
			})),
	})

	form := hb.NewForm().
		ID("FORM_TRANSACTIONS").
		Style("display: block").
		Method(http.MethodGet).
		Children([]hb.TagInterface{
			hb.NewDiv().Class("row").Children([]hb.TagInterface{
				hb.NewDiv().Class("col-md-2").Children([]hb.TagInterface{
					search,
				}),
				hb.NewDiv().Class("col-md-4").Children([]hb.TagInterface{
					period,
				}),
				hb.NewDiv().Class("col-md-2").Children([]hb.TagInterface{
					status,
				}),
				hb.NewDiv().Class("col-md-1").Children([]hb.TagInterface{
					searchButton,
				}),
			}),
		})

	return hb.NewDiv().
		Class("card bg-light mb-3").
		Style("").
		Children([]hb.TagInterface{
			hb.NewDiv().Class("card-body").Children([]hb.TagInterface{
				form,
			}),
		})
}

func (controller *blogPostManagerController) tablePagination(data blogPostManagerControllerData, count int, page int, perPage int) hb.TagInterface {
	url := links.NewAdminLinks().BlogPostManager(map[string]string{
		"search":    data.search,
		"status":    data.status,
		"date_from": data.dateFrom,
		"date_to":   data.dateTo,
		"by":        data.sortBy,
		"order":     data.sortOrder,
	})

	url = lo.Ternary(strings.Contains(url, "?"), url+"&page=", url+"?page=") // page must be last

	pagination := bs.Pagination(bs.PaginationOptions{
		NumberItems:       count,
		CurrentPageNumber: page,
		PagesToShow:       5,
		PerPage:           perPage,
		URL:               url,
	})

	return hb.NewDiv().
		Class(`d-flex justify-content-left mt-5 pagination-primary-soft rounded mb-0`).
		HTML(pagination)
}

type blogPostManagerControllerData struct {
	// r            *http.Request
	page       string
	pageInt    int
	perPage    int
	sortOrder  string
	sortBy     string
	status     string
	search     string
	customerID string
	dateFrom   string
	dateTo     string
	blogList   []blogstore.Post
	blogCount  int64
}
