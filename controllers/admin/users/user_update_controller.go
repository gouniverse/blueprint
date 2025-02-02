package admin

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"

	"github.com/asaskevich/govalidator"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/userstore"
	"github.com/gouniverse/utils"
)

// == CONTROLLER ==============================================================

type userUpdateController struct{}

var _ router.HTMLControllerInterface = (*userUpdateController)(nil)

// == CONSTRUCTOR =============================================================

func NewUserUpdateController() *userUpdateController {
	return &userUpdateController{}
}

func (controller userUpdateController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareDataAndValidate(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewAdminLinks().UsersUserManager(map[string]string{}), 10)
	}

	if r.Method == http.MethodPost {
		return controller.form(data).ToHTML()
	}

	return layouts.NewAdminLayout(r, layouts.Options{
		Title:   "Edit User | Blog",
		Content: controller.page(data),
		ScriptURLs: []string{
			cdn.Jquery_3_7_1(),
			cdn.TrumbowygJs_2_27_3(),
			cdn.Sweetalert2_10(),
			cdn.JqueryUiJs_1_13_1(), // needed for BlockArea
			links.NewWebsiteLinks().Resource(`/blockarea_v0200.js`, map[string]string{}),
		},
		Scripts: []string{
			controller.script(),
		},
		StyleURLs: []string{
			cdn.TrumbowygCss_2_27_3(),
			cdn.JqueryUiCss_1_13_1(), // needed for BlockArea
		},
	}).ToHTML()
}

func (controller userUpdateController) script() string {
	js := ``
	return js
}

func (controller userUpdateController) page(data userUpdateControllerData) hb.TagInterface {
	breadcrumbs := layouts.Breadcrumbs([]layouts.Breadcrumb{
		{
			Name: "Home",
			URL:  links.NewAdminLinks().Home(map[string]string{}),
		},
		{
			Name: "User Manager",
			URL:  links.NewAdminLinks().UsersUserManager(map[string]string{}),
		},
		{
			Name: "Edit User",
			URL:  links.NewAdminLinks().UsersUserUpdate(map[string]string{"user_id": data.userID}),
		},
	})

	buttonSave := hb.Button().
		Class("btn btn-primary ms-2 float-end").
		Child(hb.I().Class("bi bi-save").Style("margin-top:-4px;margin-right:8px;font-size:16px;")).
		HTML("Save").
		HxInclude("#FormUserUpdate").
		HxPost(links.NewAdminLinks().UsersUserUpdate(map[string]string{"userID": data.userID})).
		HxTarget("#FormUserUpdate")

	buttonCancel := hb.Hyperlink().
		Class("btn btn-secondary ms-2 float-end").
		Child(hb.I().Class("bi bi-chevron-left").Style("margin-top:-4px;margin-right:8px;font-size:16px;")).
		HTML("Back").
		Href(links.NewAdminLinks().UsersUserManager(map[string]string{}))

	heading := hb.Heading1().
		HTML("Edit User").
		Child(buttonSave).
		Child(buttonCancel)

	card := hb.Div().
		Class("card").
		Child(
			hb.Div().
				Class("card-header").
				Style(`display:flex;justify-content:space-between;align-items:center;`).
				Child(hb.Heading4().
					HTML("User Details").
					Style("margin-bottom:0;display:inline-block;")).
				Child(buttonSave),
		).
		Child(
			hb.Div().
				Class("card-body").
				Child(controller.form(data)))

	userTitle := hb.Heading2().
		Class("mb-3").
		Text("User: ").
		Text(data.user.FirstName()).
		Text(" ").
		Text(data.user.LastName())

	return hb.Div().
		Class("container").
		Child(breadcrumbs).
		Child(hb.HR()).
		Child(heading).
		Child(userTitle).
		Child(card)
}

func (controller userUpdateController) form(data userUpdateControllerData) hb.TagInterface {
	fieldsDetails := []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Label: "Status",
			Name:  "user_status",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Value: data.formStatus,
			Help:  `The status of the user.`,
			Options: []form.FieldOption{
				{
					Value: "- not selected -",
					Key:   "",
				},
				{
					Value: "Active",
					Key:   userstore.USER_STATUS_ACTIVE,
				},
				{
					Value: "Unverified",
					Key:   userstore.USER_STATUS_UNVERIFIED,
				},
				{
					Value: "Inactive",
					Key:   userstore.USER_STATUS_INACTIVE,
				},
				{
					Value: "In Trash Bin",
					Key:   userstore.USER_STATUS_DELETED,
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Label: "First Name",
			Name:  "user_first_name",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Value: data.formFirstName,
			Help:  `The first name of the user.`,
		}),
		form.NewField(form.FieldOptions{
			Label: "Last Name",
			Name:  "user_last_name",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Value: data.formLastName,
			Help:  `The last name of the user.`,
		}),
		form.NewField(form.FieldOptions{
			Label: "Email",
			Name:  "user_email",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Value: data.formEmail,
			Help:  `The email address of the user.`,
		}),
		form.NewField(form.FieldOptions{
			Label: "Admin Notes",
			Name:  "user_memo",
			Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			Value: data.formMemo,
			Help:  "Admin notes for this bloguser. These notes will not be visible to the public.",
		}),
		form.NewField(form.FieldOptions{
			Label:    "User ID",
			Name:     "user_id",
			Type:     form.FORM_FIELD_TYPE_STRING,
			Value:    data.userID,
			Readonly: true,
			Help:     "The reference number (ID) of the user.",
		}),
		form.NewField(form.FieldOptions{
			Label:    "User ID",
			Name:     "user_id",
			Type:     form.FORM_FIELD_TYPE_HIDDEN,
			Value:    data.userID,
			Readonly: true,
		}),
	}

	formUserUpdate := form.NewForm(form.FormOptions{
		ID: "FormUserUpdate",
	})

	formUserUpdate.SetFields(fieldsDetails)

	if data.formErrorMessage != "" {
		formUserUpdate.AddField(form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: hb.Swal(hb.SwalOptions{Icon: "error", Text: data.formErrorMessage}).ToHTML(),
		}))
	}

	if data.formSuccessMessage != "" {
		formUserUpdate.AddField(form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: hb.Swal(hb.SwalOptions{Icon: "success", Text: data.formSuccessMessage}).ToHTML(),
		}))
	}

	return formUserUpdate.Build()
}

func (controller userUpdateController) saveUser(r *http.Request, data userUpdateControllerData) (d userUpdateControllerData, errorMessage string) {
	if config.UserStore == nil {
		return data, "User store is not configured"
	}

	data.formFirstName = utils.Req(r, "user_first_name", "")
	data.formLastName = utils.Req(r, "user_last_name", "")
	data.formEmail = utils.Req(r, "user_email", "")
	data.formMemo = utils.Req(r, "user_memo", "")
	data.formStatus = utils.Req(r, "user_status", "")

	if data.formStatus == "" {
		data.formErrorMessage = "Status is required"
		return data, ""
	}

	if data.formFirstName == "" {
		data.formErrorMessage = "First name is required"
		return data, ""
	}

	if data.formLastName == "" {
		data.formErrorMessage = "Last name is required"
		return data, ""
	}

	if data.formEmail == "" {
		data.formErrorMessage = "Email is required"
		return data, ""
	}

	if !govalidator.IsEmail(data.formEmail) {
		data.formErrorMessage = "Invalid email address"
		return data, ""
	}

	data.user.SetMemo(data.formMemo)
	data.user.SetStatus(data.formStatus)

	err := config.UserStore.UserUpdate(r.Context(), data.user)

	if err != nil {
		config.LogStore.ErrorWithContext("At userUpdateController > prepareDataAndValidate", err.Error())
		data.formErrorMessage = "System error. Saving user failed"
		return data, ""
	}

	err = userTokenize(data.user, data.formFirstName, data.formLastName, data.formEmail)

	if err != nil {
		config.LogStore.ErrorWithContext("At userUpdateController > prepareDataAndValidate", err.Error())
		data.formErrorMessage = "System error. Saving user failed"
		return data, ""
	}

	data.formSuccessMessage = "User saved successfully"

	return data, ""
}

func (controller userUpdateController) prepareDataAndValidate(r *http.Request) (data userUpdateControllerData, errorMessage string) {
	data.action = utils.Req(r, "action", "")
	data.userID = utils.Req(r, "user_id", "")

	if data.userID == "" {
		return data, "User ID is required"
	}

	user, err := config.UserStore.UserFindByID(r.Context(), data.userID)

	if err != nil {
		config.LogStore.ErrorWithContext("At userUpdateController > prepareDataAndValidate", err.Error())
		return data, "User not found"
	}

	if user == nil {
		return data, "User not found"
	}

	data.user = user

	firstName, lastName, email, err := helpers.UserUntokenized(r.Context(), data.user)

	if err != nil {
		config.LogStore.ErrorWithContext("At userManagerController > tableUsers", err.Error())
		return data, "Tokens failed to be read"
	}

	data.formFirstName = firstName
	data.formLastName = lastName
	data.formEmail = email
	data.formMemo = data.user.Memo()
	data.formStatus = data.user.Status()

	if r.Method != http.MethodPost {
		return data, ""
	}

	return controller.saveUser(r, data)
}

type userUpdateControllerData struct {
	action string
	userID string
	user   userstore.UserInterface

	formErrorMessage   string
	formSuccessMessage string
	formEmail          string
	formFirstName      string
	formLastName       string
	formMemo           string
	formStatus         string
}
