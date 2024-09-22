package admin

import (
	"net/http"
	"project/config"
	"project/pkg/userstore"

	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/blindindexstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

const ActionModalUserFilterShow = "modal_user_filter_show"

// == CONTROLLER ==============================================================

type userManagerController struct{}

var _ router.HTMLControllerInterface = (*userManagerController)(nil)

// == CONSTRUCTOR =============================================================

func NewUserManagerController() *userManagerController {
	return &userManagerController{}
}

func (controller *userManagerController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareData(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewAdminLinks().Home(map[string]string{}), 10)
	}

	if data.action == ActionModalUserFilterShow {
		return controller.onModalUserFilterShow(data).ToHTML()
	}

	return layouts.NewAdminLayout(r, layouts.Options{
		Title:   "Users | User Manager",
		Content: controller.page(data),
		ScriptURLs: []string{
			cdn.Htmx_1_9_4(),
			cdn.Sweetalert2_10(),
		},
		Styles: []string{},
	}).ToHTML()
}

func (controller *userManagerController) onModalUserFilterShow(data userManagerControllerData) *hb.Tag {
	modalCloseScript := `document.getElementById('ModalMessage').remove();document.getElementById('ModalBackdrop').remove();`

	title := hb.NewHeading5().
		Text("Filters").
		Style(`margin:0px;padding:0px;`)

	buttonModalClose := hb.NewButton().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	buttonCancel := hb.NewButton().
		Child(hb.NewI().Class("bi bi-chevron-left me-2")).
		HTML("Cancel").
		Class("btn btn-secondary float-start").
		OnClick(modalCloseScript)

	buttonOk := hb.NewButton().
		Child(hb.NewI().Class("bi bi-check me-2")).
		HTML("Apply").
		Class("btn btn-primary float-end").
		OnClick(`FormFilters.submit();` + modalCloseScript)

	filterForm := form.NewForm(form.FormOptions{
		ID:     "FormFilters",
		Method: http.MethodGet,
		Fields: []form.Field{
			{
				Label: "Status",
				Name:  "status",
				Type:  form.FORM_FIELD_TYPE_SELECT,
				Help:  `The status of the user.`,
				Value: data.formStatus,
				Options: []form.FieldOption{
					{
						Value: "",
						Key:   "",
					},
					{
						Value: "Active",
						Key:   userstore.USER_STATUS_ACTIVE,
					},
					{
						Value: "Inactive",
						Key:   userstore.USER_STATUS_INACTIVE,
					},
					{
						Value: "Unverified",
						Key:   userstore.USER_STATUS_UNVERIFIED,
					},
					{
						Value: "Deleted",
						Key:   userstore.USER_STATUS_DELETED,
					},
				},
			},
			{
				Label: "First Name",
				Name:  "first_name",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formFirstName,
				Help:  `Filter by first name.`,
			},
			{
				Label: "Last Name",
				Name:  "last_name",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formLastName,
				Help:  `Filter by last name.`,
			},
			{
				Label: "Email",
				Name:  "email",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formEmail,
				Help:  `Filter by email.`,
			},
			{
				Label: "Created From",
				Name:  "created_from",
				Type:  form.FORM_FIELD_TYPE_DATE,
				Value: data.formCreatedFrom,
				Help:  `Filter by creation date.`,
			},
			{
				Label: "Created To",
				Name:  "created_to",
				Type:  form.FORM_FIELD_TYPE_DATE,
				Value: data.formCreatedTo,
				Help:  `Filter by creation date.`,
			},
			{
				Label: "User ID",
				Name:  "user_id",
				Type:  form.FORM_FIELD_TYPE_STRING,
				Value: data.formUserID,
				Help:  `Find user by reference number (ID).`,
			},
		},
	}).Build()

	modal := bs.Modal().
		ID("ModalMessage").
		Class("fade show").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Children([]hb.TagInterface{
			bs.ModalDialog().Children([]hb.TagInterface{
				bs.ModalContent().Children([]hb.TagInterface{
					bs.ModalHeader().Children([]hb.TagInterface{
						title,
						buttonModalClose,
					}),

					bs.ModalBody().
						Child(filterForm),

					bs.ModalFooter().
						Style(`display:flex;justify-content:space-between;`).
						Child(buttonCancel).
						Child(buttonOk),
				}),
			}),
		})

	backdrop := hb.NewDiv().
		ID("ModalBackdrop").
		Class("modal-backdrop fade show").
		Style("display:block;")

	return hb.NewWrap().Children([]hb.TagInterface{
		modal,
		backdrop,
	})

}

func (controller *userManagerController) page(data userManagerControllerData) hb.TagInterface {
	breadcrumbs := layouts.Breadcrumbs([]layouts.Breadcrumb{
		{
			Name: "Home",
			URL:  links.NewAdminLinks().Home(map[string]string{}),
		},
		{
			Name: "Users",
			URL:  links.NewAdminLinks().UsersUserManager(map[string]string{}),
		},
		{
			Name: "User Manager",
			URL:  links.NewAdminLinks().UsersUserManager(map[string]string{}),
		},
	})

	buttonUserNew := hb.NewButton().
		Class("btn btn-primary float-end").
		Child(hb.NewI().Class("bi bi-plus-circle").Style("margin-top:-4px;margin-right:8px;font-size:16px;")).
		HTML("New User").
		HxGet(links.NewAdminLinks().UsersUserCreate(map[string]string{})).
		HxTarget("body").
		HxSwap("beforeend")

	title := hb.NewHeading1().
		HTML("Users. User Manager").
		Child(buttonUserNew)

	return hb.NewDiv().
		Class("container").
		Child(breadcrumbs).
		Child(hb.NewHR()).
		Child(title).
		Child(controller.tableUsers(data))
}

func (controller *userManagerController) tableUsers(data userManagerControllerData) hb.TagInterface {
	table := hb.NewTable().
		Class("table table-striped table-hover table-bordered").
		Children([]hb.TagInterface{
			hb.NewThead().Children([]hb.TagInterface{
				hb.NewTR().Children([]hb.TagInterface{
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "First Name", "first_name")).
						Text(", ").
						Child(controller.sortableColumnLabel(data, "Last Name", "last_name")).
						Text(", ").
						Child(controller.sortableColumnLabel(data, "Reference", "id")).
						Style(`cursor: pointer;`),
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "Status", "status")).
						Style("width: 200px;cursor: pointer;"),
					hb.NewTH().
						Child(controller.sortableColumnLabel(data, "E-mail", "email")).
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
			hb.NewTbody().Children(lo.Map(data.userList, func(user userstore.User, _ int) hb.TagInterface {
				firstName, lastName, email, err := helpers.UserUntokenized(user)

				if err != nil {
					config.LogStore.ErrorWithContext("At userManagerController > tableUsers", err.Error())
					firstName = "n/a"
					lastName = "n/a"
					email = "n/a"
				}

				userLink := hb.NewHyperlink().
					Text(firstName).
					Text(` `).
					Text(lastName).
					Href(links.NewAdminLinks().UsersUserUpdate(map[string]string{"user_id": user.ID()}))

				status := hb.NewSpan().
					Style(`font-weight: bold;`).
					StyleIf(user.IsActive(), `color:green;`).
					StyleIf(user.IsDeleted(), `color:silver;`).
					StyleIf(user.IsUnverified(), `color:blue;`).
					StyleIf(user.IsInactive(), `color:red;`).
					HTML(user.Status())

				buttonEdit := hb.NewHyperlink().
					Class("btn btn-primary me-2").
					Child(hb.NewI().Class("bi bi-pencil-square")).
					Title("Edit").
					Href(links.NewAdminLinks().UsersUserUpdate(map[string]string{"user_id": user.ID()})).
					Target("_blank")

				buttonDelete := hb.NewHyperlink().
					Class("btn btn-danger").
					Child(hb.NewI().Class("bi bi-trash")).
					Title("Delete").
					HxGet(links.NewAdminLinks().UsersUserDelete(map[string]string{"user_id": user.ID()})).
					HxTarget("body").
					HxSwap("beforeend")

				buttonImpersonate := hb.NewHyperlink().
					Class("btn btn-warning me-2").
					Child(hb.NewI().Class("bi bi-shuffle")).
					Title("Impersonate").
					Href(links.NewAdminLinks().UsersUserImpersonate(map[string]string{"user_id": user.ID()}))

				return hb.NewTR().Children([]hb.TagInterface{
					hb.NewTD().
						Child(hb.NewDiv().Child(userLink)).
						Child(hb.NewDiv().
							Style("font-size: 11px;").
							HTML("Ref: ").
							HTML(user.ID())),
					hb.NewTD().
						Child(status),
					hb.NewTD().
						Child(hb.NewDiv().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(email)),
					hb.NewTD().
						Child(hb.NewDiv().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(user.CreatedAtCarbon().Format("d M Y"))),
					hb.NewTD().
						Child(hb.NewDiv().
							Style("font-size: 13px;white-space: nowrap;").
							HTML(user.UpdatedAtCarbon().Format("d M Y"))),
					hb.NewTD().
						Child(buttonEdit).
						Child(buttonImpersonate).
						Child(buttonDelete),
				})
			})),
		})

	// cfmt.Successln("Table: ", table)

	return hb.NewWrap().Children([]hb.TagInterface{
		controller.tableFilter(data),
		table,
		controller.tablePagination(data, int(data.userCount), data.pageInt, data.perPage),
	})
}

func (controller *userManagerController) sortableColumnLabel(data userManagerControllerData, tableLabel string, columnName string) hb.TagInterface {
	isSelected := strings.EqualFold(data.sortBy, columnName)

	direction := lo.If(data.sortOrder == "asc", "desc").Else("asc")

	if !isSelected {
		direction = "asc"
	}

	link := links.NewAdminLinks().UsersUserManager(map[string]string{
		"page":      "0",
		"by":        columnName,
		"sort":      direction,
		"date_from": data.formCreatedFrom,
		"date_to":   data.formCreatedTo,
		"status":    data.formStatus,
		"user_id":   data.formUserID,
	})
	return hb.NewHyperlink().
		HTML(tableLabel).
		Child(controller.sortingIndicator(columnName, data.sortBy, direction)).
		Href(link)
}

func (controller *userManagerController) sortingIndicator(columnName string, sortByColumnName string, sortOrder string) hb.TagInterface {
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

func (controller *userManagerController) tableFilter(data userManagerControllerData) hb.TagInterface {
	buttonFilter := hb.NewButton().
		Class("btn btn-sm btn-info me-2").
		Style("margin-bottom: 2px; margin-left:2px; margin-right:2px;").
		Child(hb.NewI().Class("bi bi-filter me-2")).
		Text("Filters").
		HxPost(links.NewAdminLinks().UsersUserManager(map[string]string{
			"action":       ActionModalUserFilterShow,
			"first_name":   data.formFirstName,
			"last_name":    data.formLastName,
			"email":        data.formEmail,
			"status":       data.formStatus,
			"user_id":      data.formUserID,
			"created_from": data.formCreatedFrom,
			"created_to":   data.formCreatedTo,
		})).
		HxTarget("body").
		HxSwap("beforeend")

	description := []string{
		hb.NewSpan().HTML("Showing users").Text(" ").ToHTML(),
	}

	if data.formStatus != "" {
		description = append(description, hb.NewSpan().Text("with status: "+data.formStatus).ToHTML())
	} else {
		description = append(description, hb.NewSpan().Text("with status: any").ToHTML())
	}

	if data.formEmail != "" {
		description = append(description, hb.NewSpan().Text("and email: "+data.formEmail).ToHTML())
	}

	if data.formUserID != "" {
		description = append(description, hb.NewSpan().Text("and ID: "+data.formUserID).ToHTML())
	}

	if data.formFirstName != "" {
		description = append(description, hb.NewSpan().Text("and first name: "+data.formFirstName).ToHTML())
	}

	if data.formLastName != "" {
		description = append(description, hb.NewSpan().Text("and last name: "+data.formLastName).ToHTML())
	}

	if data.formCreatedFrom != "" && data.formCreatedTo != "" {
		description = append(description, hb.NewSpan().Text("and created between: "+data.formCreatedFrom+" and "+data.formCreatedTo).ToHTML())
	} else if data.formCreatedFrom != "" {
		description = append(description, hb.NewSpan().Text("and created after: "+data.formCreatedFrom).ToHTML())
	} else if data.formCreatedTo != "" {
		description = append(description, hb.NewSpan().Text("and created before: "+data.formCreatedTo).ToHTML())
	}

	return hb.NewDiv().
		Class("card bg-light mb-3").
		Style("").
		Children([]hb.TagInterface{
			hb.NewDiv().Class("card-body").
				Child(buttonFilter).
				Child(hb.NewSpan().
					HTML(strings.Join(description, " "))),
		})
}

func (controller *userManagerController) tablePagination(data userManagerControllerData, count int, page int, perPage int) hb.TagInterface {
	url := links.NewAdminLinks().UsersUserManager(map[string]string{
		"status":       data.formStatus,
		"first_name":   data.formFirstName,
		"last_name":    data.formLastName,
		"email":        data.formEmail,
		"created_from": data.formCreatedFrom,
		"created_to":   data.formCreatedTo,
		"by":           data.sortBy,
		"order":        data.sortOrder,
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

func (controller *userManagerController) prepareData(r *http.Request) (data userManagerControllerData, errorMessage string) {
	var err error
	data.request = r
	data.action = utils.Req(r, "action", "")
	data.page = utils.Req(r, "page", "0")
	data.pageInt = cast.ToInt(data.page)
	data.perPage = cast.ToInt(utils.Req(r, "per_page", "10"))
	data.sortOrder = utils.Req(r, "sort_order", sb.DESC)
	data.sortBy = utils.Req(r, "by", userstore.COLUMN_CREATED_AT)
	data.formEmail = utils.Req(r, "email", "")
	data.formFirstName = utils.Req(r, "first_name", "")
	data.formLastName = utils.Req(r, "last_name", "")
	data.formStatus = utils.Req(r, "status", "")
	data.formCreatedFrom = utils.Req(r, "created_from", "")
	data.formCreatedTo = utils.Req(r, "created_to", "")

	userList, userCount, err := controller.fetchUserList(data)

	if err != nil {
		config.LogStore.ErrorWithContext("At userManagerController > prepareData", err.Error())
		return data, "error retrieving users"
	}

	data.userList = userList
	data.userCount = userCount

	return data, ""
}

func (controller *userManagerController) fetchUserList(data userManagerControllerData) ([]userstore.User, int64, error) {
	userIDs := []string{}

	if data.formFirstName != "" {
		firstNameUserIDs, err := config.BlindIndexStoreFirstName.Search(data.formFirstName, blindindexstore.SEARCH_TYPE_CONTAINS)

		if err != nil {
			config.LogStore.ErrorWithContext("At userManagerController > prepareData", err.Error())
			return []userstore.User{}, 0, err
		}

		if len(firstNameUserIDs) == 0 {
			return []userstore.User{}, 0, nil
		}

		userIDs = append(userIDs, firstNameUserIDs...)
	}

	if data.formLastName != "" {
		lastNameUserIDs, err := config.BlindIndexStoreLastName.Search(data.formLastName, blindindexstore.SEARCH_TYPE_CONTAINS)

		if err != nil {
			config.LogStore.ErrorWithContext("At userManagerController > prepareData", err.Error())
			return []userstore.User{}, 0, err
		}

		if len(lastNameUserIDs) == 0 {
			return []userstore.User{}, 0, nil
		}

		userIDs = append(userIDs, lastNameUserIDs...)
	}

	if data.formEmail != "" {
		emailUserIDs, err := config.BlindIndexStoreEmail.Search(data.formEmail, blindindexstore.SEARCH_TYPE_CONTAINS)

		if err != nil {
			config.LogStore.ErrorWithContext("At userManagerController > prepareData", err.Error())
			return []userstore.User{}, 0, err
		}

		if len(emailUserIDs) == 0 {
			return []userstore.User{}, 0, nil
		}

		userIDs = append(userIDs, emailUserIDs...)
	}

	query := userstore.UserQueryOptions{
		IDIn:      userIDs,
		Offset:    data.pageInt * data.perPage,
		Limit:     data.perPage,
		Status:    data.formStatus,
		SortOrder: data.sortOrder,
		OrderBy:   data.sortBy,
	}

	if data.formCreatedFrom != "" {
		query.CreatedAtGte = data.formCreatedFrom + " 00:00:00"
	}

	if data.formCreatedTo != "" {
		query.CreatedAtLte = data.formCreatedTo + " 23:59:59"
	}

	userList, err := config.UserStore.UserList(query)

	if err != nil {
		config.LogStore.ErrorWithContext("At userManagerController > prepareData", err.Error())
		return []userstore.User{}, 0, err
	}

	userCount, err := config.UserStore.UserCount(query)

	if err != nil {
		config.LogStore.ErrorWithContext("At userManagerController > prepareData", err.Error())
		return []userstore.User{}, 0, err
	}

	return userList, userCount, nil
}

type userManagerControllerData struct {
	request         *http.Request
	action          string
	page            string
	pageInt         int
	perPage         int
	sortOrder       string
	sortBy          string
	formStatus      string
	formEmail       string
	formFirstName   string
	formLastName    string
	formCreatedFrom string
	formCreatedTo   string
	formUserID      string
	userList        []userstore.User
	userCount       int64
}
