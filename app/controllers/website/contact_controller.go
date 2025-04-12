package website

import (
	"context"
	"net/http"
	"project/app/layouts"
	"project/app/links"
	"project/app/tasks"
	"project/config"
	"project/internal/helpers"
	"strings"

	"github.com/gouniverse/csrf"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/customstore"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

// == CONTROLLER ==============================================================

type contactController struct{}

type contactControllerData struct {
	email          string
	firstName      string
	lastName       string
	text           string
	csrfToken      string
	errorMessage   string
	successMessage string
	redirectURL    string
}

// == CONSTRUCTOR =============================================================

func NewContactController() *contactController {
	return &contactController{}
}

func (controller *contactController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	authUser := helpers.GetAuthUser(r)
	data := contactControllerData{
		email:     lo.TernaryF(authUser != nil, func() string { return authUser.Email() }, func() string { return "" }),
		firstName: lo.TernaryF(authUser != nil, func() string { return authUser.FirstName() }, func() string { return "" }),
		lastName:  lo.TernaryF(authUser != nil, func() string { return authUser.LastName() }, func() string { return "" }),
		text:      "",
		csrfToken: csrf.TokenGenerate("holymoly"),
	}

	title := hb.Heading1().HTML("Contact").Style("margin:30px 0px 30px 0px;")

	paragraph1 := hb.Paragraph().
		HTML("Please add and check your details below are correct so that we can respond to you as requested.").
		Style("margin-bottom:20px;")

	formContact := controller.contactForm(r, data)

	page := hb.Section().
		Child(
			hb.Div().
				Class("container").
				Child(title).
				Child(paragraph1).
				Child(formContact),
		)

	return layouts.NewUserLayout(r, layouts.Options{
		Title:      "Contact",
		Content:    page,
		ScriptURLs: []string{cdn.Sweetalert2_11()},
	}).ToHTML()
}

func (controller *contactController) PostSubmit(w http.ResponseWriter, r *http.Request) string {
	data := contactControllerData{
		email:     strings.TrimSpace(utils.Req(r, "email", "")),
		firstName: strings.TrimSpace(utils.Req(r, "first_name", "")),
		lastName:  strings.TrimSpace(utils.Req(r, "last_name", "")),
		text:      strings.TrimSpace(utils.Req(r, "text", "")),
		csrfToken: strings.TrimSpace(utils.Req(r, "csrf_token", "")),
	}

	if data.csrfToken == "" {
		data.errorMessage = "CSRF token is required"
		data.redirectURL = links.NewWebsiteLinks().Contact()
		return controller.contactForm(r, data).ToHTML()
	}

	if !csrf.TokenValidate(data.csrfToken, "holymoly") {
		data.errorMessage = "CSRF token is invalid"
		data.redirectURL = links.NewWebsiteLinks().Contact()
		return controller.contactForm(r, data).ToHTML()
	}

	if data.firstName == "" {
		data.errorMessage = "First name is required"
		return controller.contactForm(r, data).ToHTML()
	}

	if data.lastName == "" {
		data.errorMessage = "Last name is required"
		return controller.contactForm(r, data).ToHTML()
	}

	if data.email == "" {
		data.errorMessage = "Email is required"
		return controller.contactForm(r, data).ToHTML()
	}

	if data.text == "" {
		data.errorMessage = "Text is required"
		return controller.contactForm(r, data).ToHTML()
	}

	record := customstore.NewRecord("contact")

	record.SetPayloadMap(map[string]interface{}{
		"first_name": data.firstName,
		"last_name":  data.lastName,
		"email":      data.email,
		"text":       data.text,
	})

	err := config.CustomStore.RecordCreate(record)

	if err != nil {
		config.Logger.Error("At contactController.PostSubmit", "error", err.Error())
		data.errorMessage = "System error occurred. Please try again later."
		return controller.contactForm(r, data).ToHTML()
	}

	_, err = tasks.NewEmailToAdminOnNewContactFormSubmittedTaskHandler().Enqueue()

	if err != nil {
		config.Logger.Error("At contactController.PostSubmit. Enqueue EmailToAdminOnNewContactFormSubmittedTask", "error", err.Error())
		// No need to return error here
	}

	authUser := helpers.GetAuthUser(r)
	canUpdateFirstName, canUpdateLastName, _ := controller.canUpdateFirstNameLastNameEmail(r)

	if authUser != nil && (canUpdateFirstName || canUpdateLastName) && config.UserStore != nil {
		authUser.SetFirstName(data.firstName)
		authUser.SetLastName(data.lastName)

		err = config.UserStore.UserUpdate(context.Background(), authUser)

		if err != nil {
			config.Logger.Error("At contactController.PostSubmit", "error", err.Error())
			// No need to return error here
		}
	}

	data.successMessage = "Your message has been sent."
	data.redirectURL = links.NewWebsiteLinks().Contact()
	return controller.contactForm(r, data).ToHTML()
}

// canUpdateFirstNameLastNameEmail returns a boolean value for each field
// indicating if it can be updated or not.
//
// Business Logic:
// - If user is not logged in, then all fields can be updated.
// - If first name is not already set, then it can be updated.
// - If last name is not already set, then it can be updated.
// - If email is not already set, then it can be updated.
//
// Params:
// r: The http.Request object.
//
// Returns:
// - canUpdateFirstName: Whether the first name can be updated or not.
// - canUpdateLastName: Whether the last name can be updated or not.
// - canUpdateEmail: Whether the email can be updated or not.
func (c *contactController) canUpdateFirstNameLastNameEmail(r *http.Request) (canUpdateFirstName, canUpdateLastName, canUpdateEmail bool) {
	authUser := helpers.GetAuthUser(r)

	if authUser == nil {
		return true, true, true
	}

	canUpdateFirstName = lo.Ternary(authUser.FirstName() == "", true, false)
	canUpdateLastName = lo.Ternary(authUser.LastName() == "", true, false)
	canUpdateEmail = lo.Ternary(authUser.Email() == "", true, false)

	return canUpdateFirstName, canUpdateLastName, canUpdateEmail
}

// contactForm builds the contact form.
//
// Params:
//   - r: The http.Request object.
//   - data: The data for the contact form.
//
// Returns:
// - *hb.Tag: The contact form.
func (controller *contactController) contactForm(r *http.Request, data contactControllerData) *hb.Tag {
	canUpdateFirstName, canUpdateLastName, canUpdateEmail := controller.canUpdateFirstNameLastNameEmail(r)

	required := hb.Sup().
		HTML("required").
		Style("margin-left:5px;color:lightcoral;")

	firstName := bs.FormGroup().Children([]hb.TagInterface{
		bs.FormLabel("First name").
			Child(required),
		bs.FormInput().
			Name("first_name").
			Value(data.firstName).
			AttrIf(!canUpdateFirstName, "readonly", "readonly").
			StyleIf(!canUpdateFirstName, "background-color:#ccc;"),
	})

	lastName := bs.FormGroup().Children([]hb.TagInterface{
		bs.FormLabel("Last name").Child(required),
		bs.FormInput().
			Name("last_name").
			Value(data.lastName).
			AttrIf(!canUpdateLastName, "readonly", "readonly").
			StyleIf(!canUpdateLastName, "background-color:#ccc;"),
	})

	email := bs.FormGroup().Children([]hb.TagInterface{
		bs.FormLabel("Email").
			Child(required),
		bs.FormInput().
			Name("email").
			Value(data.email).
			AttrIf(!canUpdateEmail, "readonly", "readonly").
			StyleIf(!canUpdateEmail, "background-color:#ccc;"),
	})

	csrfToken := bs.FormGroup().Children([]hb.TagInterface{
		bs.FormInput().Type(hb.TYPE_HIDDEN).
			Name("csrf_token").
			Value(data.csrfToken),
	})

	text := bs.FormGroup().Children([]hb.TagInterface{
		bs.FormLabel("Text").
			Child(required),
		bs.FormTextArea().
			Name("text").
			Value(data.text).
			HTML(data.text).
			Style("height:200px;"),
	})

	buttonSubmit := bs.Button().
		Class("btn-primary mb-0").
		Attr("type", "button").
		Child(hb.I().Class("bi bi-rocket me-2")).
		HTML("Send ").
		HxInclude("#FormContact").
		HxTarget("#CardContact").
		HxTrigger("click").
		HxSwap("outerHTML").
		HxPost(links.NewWebsiteLinks().Contact())

	formContact := hb.Div().
		ID("FormContact").
		Children([]hb.TagInterface{
			hb.Div().
				Class("row g-4").
				// First name
				Child(hb.Div().
					Class("col-6").
					Child(firstName)).
				// Last name
				Child(hb.Div().
					Class("col-6").
					Child(lastName)).
				// Email
				Child(
					hb.Div().
						Class("col-12").
						Child(email)).
				// Text
				Child(hb.Div().
					Class("col-12").
					Child(text)),
			hb.Div().
				Class("row mt-3").
				// Button Submit
				Child(hb.Div().
					Class("col-12").
					Class("d-sm-flex justify-content-end").
					Child(buttonSubmit)),
		}).Child(csrfToken)

	errorMessageJSON, _ := utils.ToJSON(data.errorMessage)
	successMessageJSON, _ := utils.ToJSON(data.successMessage)
	return hb.Div().ID("CardContact").
		Class("card bg-transparent border rounded-3").
		Style("text-align:left;").
		Children([]hb.TagInterface{
			// hb.Div().Class("card-header  bg-transparent").Children([]hb.TagInterface{
			// 	hb.Heading3().
			// 		HTML("User Details").
			// 		Style("text-align:left;font-size:23px;color:#333;"),
			// }),
			hb.Div().Class("card-body").Children([]hb.TagInterface{
				formContact,
			}),
		}).
		ChildIf(data.errorMessage != "", hb.Script(`
			Swal.fire({
				icon: 'error',
				title: 'Oops...',
				text: `+errorMessageJSON+`,
			})
		`)).
		ChildIf(data.successMessage != "", hb.Script(`
			Swal.fire({
				icon: 'success',
				title: 'Saved',
				text: `+successMessageJSON+`,
			})
		`)).
		ChildIf(data.redirectURL != "", hb.Script(`setTimeout(() => {window.location.href="`+data.redirectURL+`";}, 5000);`))
}
