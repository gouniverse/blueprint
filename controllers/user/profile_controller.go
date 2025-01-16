package user

import (
	"context"
	"net/http"
	"project/config"
	"project/controllers/user/partials"
	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/userstore"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

// == CONTROLLER ==============================================================

type profileController struct {
	actionOnCountrySelectedTimezoneOptions string
	formCountry                            string
	formTimezone                           string
}

var _ router.HTMLControllerInterface = (*profileController)(nil)

// == CONSTRUCTOR =============================================================

func NewProfileController() *profileController {
	return &profileController{
		actionOnCountrySelectedTimezoneOptions: "on-country-selected-timezone-options",
		formCountry:                            "country",
		formTimezone:                           "timezone",
	}
}

// == PUBLIC METHODS ==========================================================

func (controller *profileController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareData(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewUserLinks().Home(map[string]string{}), 10)
	}

	if data.action == controller.actionOnCountrySelectedTimezoneOptions {
		return controller.onCountrySelectedTimezoneOptions(data)
	}

	if r.Method == http.MethodPost {
		return controller.postUpdate(data)
	}

	breadcrumbs := layouts.NewUserBreadcrumbsSectionWithContainer([]bs.Breadcrumb{
		{
			Name: "My Profile",
			URL:  links.NewUserLinks().Profile(map[string]string{}),
		},
	})

	title := hb.Heading1().
		Text("My Account").
		Style("margin:30px 0px 30px 0px;")

	paragraph1 := hb.Paragraph().
		Text("Please keep your details updated so that we can contact you if you need our help.").
		Style("margin-bottom:20px;")

	formProfile := controller.formProfile(data)

	page := hb.Section().
		Child(breadcrumbs).
		Child(hb.HR()).
		Child(partials.UserQuickLinks(data.request)).
		Child(hb.HR()).
		Child(
			hb.Div().
				Class("container").
				Child(title).
				Child(paragraph1).
				Child(formProfile).
				Child(hb.BR()).
				Child(hb.BR()),
		)

	return layouts.NewUserLayout(r, layouts.Options{
		Title:      "My Profile",
		Content:    page,
		ScriptURLs: []string{cdn.Sweetalert2_10()},
	}).ToHTML()
}

func (controller *profileController) postUpdate(data profileControllerData) string {
	if data.firstName == "" {
		data.formErrorMessage = "First name is required field"
		return controller.formProfile(data).ToHTML()
	}

	if data.lastName == "" {
		data.formErrorMessage = "Last name is required field"
		return controller.formProfile(data).ToHTML()
	}

	if data.email == "" {
		data.formErrorMessage = "Email is required field"
		return controller.formProfile(data).ToHTML()
	}

	if data.country == "" {
		data.formErrorMessage = "Country is required field"
		return controller.formProfile(data).ToHTML()
	}

	if data.timezone == "" {
		data.formErrorMessage = "Timezone is required field"
		return controller.formProfile(data).ToHTML()
	}

	// First name
	err := config.VaultStore.TokenUpdate(data.request.Context(), data.authUser.FirstName(), data.firstName, config.VaultKey)

	if err != nil {
		config.Logger.Error("Error saving first name", "error", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	// Last name
	err = config.VaultStore.TokenUpdate(data.request.Context(), data.authUser.LastName(), data.lastName, config.VaultKey)

	if err != nil {
		config.Logger.Error("Error saving last name", "error", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	// Business name
	err = config.VaultStore.TokenUpdate(data.request.Context(), data.authUser.BusinessName(), data.buinessName, config.VaultKey)

	if err != nil {
		config.Logger.Error("Error saving business name", "error", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	// Phone
	err = config.VaultStore.TokenUpdate(data.request.Context(), data.authUser.Phone(), data.phone, config.VaultKey)

	if err != nil {
		config.Logger.Error("Error saving phone", "error", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	data.authUser.SetCountry(data.country)
	data.authUser.SetTimezone(data.timezone)

	if config.UserStore == nil {
		config.Logger.Warn("At profileController > post update. UserStore is nil.")
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	err = config.UserStore.UserUpdate(context.Background(), data.authUser)

	if err != nil {
		config.Logger.Error("Error updating user profile", "error", err.Error())

		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	data.formSuccessMessage = "Profile updated successfully"
	data.formRedirectURL = helpers.ToFlashSuccessURL(data.formSuccessMessage, links.NewUserLinks().Home(map[string]string{}), 5)
	return controller.formProfile(data).ToHTML()
}

func (controller *profileController) formProfile(data profileControllerData) *hb.Tag {
	required := hb.Sup().
		Text("required").
		Style("margin-left:5px;color:lightcoral;")

	groupFirstName := bs.FormGroup().
		Child(bs.FormLabel("First name").
			Child(required)).
		Child(bs.FormInput().
			Name("first_name").
			Value(data.firstName))

	groupLastName := bs.FormGroup().
		Child(bs.FormLabel("Last name").
			Child(required)).
		Child(bs.FormInput().
			Name("last_name").
			Value(data.lastName))

	groupEmail := bs.FormGroup().
		Child(bs.FormLabel("Email").
			Child(required)).
		Child(bs.FormInput().
			Name("email").
			Value(data.email).
			Attr("readonly", "readonly").
			Style("background-color:#F8F8F8;"))

	groupBuinessName := bs.FormGroup().
		Child(bs.FormLabel("Company / buiness name")).
		Child(bs.FormInput().
			Name("business_name").
			Value(data.buinessName))

	groupPhone := bs.FormGroup().
		Child(bs.FormLabel("Phone")).
		Child(bs.FormInput().
			Name("phone").
			Value(data.phone))

	selectCountries := bs.FormSelect().
		ID("SelectCountries").
		Name(controller.formCountry).
		Child(bs.FormSelectOption("", "")).
		Children(lo.Map(data.countryList, func(country geostore.Country, _ int) hb.TagInterface {
			return bs.FormSelectOption(country.IsoCode2(), country.Name()).
				AttrIf(data.country == country.IsoCode2(), "selected", "selected")
		})).
		HxPost(links.NewAuthLinks().Register(map[string]string{
			"action": controller.actionOnCountrySelectedTimezoneOptions,
		})).
		HxTarget("#SelectTimezones").
		HxSwap("outerHTML")

	countryGroup := hb.Div().
		Class("form-group").
		Children([]hb.TagInterface{
			bs.FormLabel("Country").
				Child(required),
			selectCountries,
		})

	timezoneGroup := hb.Div().
		Class("form-group").
		Children([]hb.TagInterface{
			bs.FormLabel("Timezone").
				Child(required),
			controller.selectTimezoneByCountry(data.country, data.timezone),
		})

	buttonSave := bs.Button().
		Class("btn-primary mb-0").
		Attr("type", "button").
		Text("Save changes").
		HxInclude("#FormProfile").
		HxTarget("#CardUserProfile").
		HxTrigger("click").
		HxSwap("outerHTML").
		HxPost(links.NewUserLinks().Profile(map[string]string{}))

	formProfile := hb.Div().ID("FormProfile").Children([]hb.TagInterface{
		bs.Row().
			Class("g-4").
			Children([]hb.TagInterface{
				bs.Column(12).Child(groupEmail),
				bs.Column(6).Child(groupFirstName),
				bs.Column(6).Child(groupLastName),
				bs.Column(6).Child(groupBuinessName),
				bs.Column(6).Child(groupPhone),
				bs.Column(6).Child(countryGroup),
				bs.Column(6).Child(timezoneGroup),
			}),
		bs.Row().
			Class("mt-3").
			Child(
				bs.Column(12).
					Class("d-sm-flex justify-content-end").
					Child(buttonSave),
			),
	})

	return hb.Div().ID("CardUserProfile").
		Class("card bg-transparent border rounded-3").
		Style("text-align:left;").
		Children([]hb.TagInterface{
			hb.Div().Class("card-header  bg-transparent").Children([]hb.TagInterface{
				hb.Heading3().
					Text("Your Details").
					Style("text-align:left;font-size:23px;color:#333;"),
			}),
			hb.Div().Class("card-body").Children([]hb.TagInterface{
				formProfile,
			}),
		}).
		ChildIf(data.formErrorMessage != "", hb.Swal(hb.SwalOptions{
			Icon:              "error",
			Title:             "Error",
			Text:              data.formErrorMessage,
			ShowCancelButton:  false,
			ConfirmButtonText: "OK",
		})).
		ChildIf(data.formSuccessMessage != "", hb.Swal(hb.SwalOptions{
			Icon:              "success",
			Title:             "Saved",
			Text:              data.formSuccessMessage,
			ShowCancelButton:  false,
			ConfirmButtonText: "OK",
			ConfirmCallback:   "window.location.href = window.location.href",
		})).
		ChildIf(data.formRedirectURL != "", hb.Script(`window.location.href = '`+data.formRedirectURL+`'`))

}

func (controller *profileController) onCountrySelectedTimezoneOptions(data profileControllerData) string {
	return controller.selectTimezoneByCountry(data.country, data.timezone).ToHTML()
}

func (controller *profileController) selectTimezoneByCountry(country string, selectedTimezone string) *hb.Tag {
	query := geostore.TimezoneQueryOptions{
		SortOrder: sb.ASC,
		OrderBy:   geostore.COLUMN_TIMEZONE,
	}

	if country != "" {
		query.CountryCode = country
	}

	timezones, errZones := config.GeoStore.TimezoneList(query)

	if errZones != nil {
		config.Logger.Error("Error listing timezones", "error", errZones.Error())
		return hb.Text("Error listing timezones")
	}

	selectTimezones := bs.FormSelect().
		ID("SelectTimezones").
		Name(controller.formTimezone).
		Child(bs.FormSelectOption("", "")).
		Children(lo.Map(timezones, func(timezone geostore.Timezone, _ int) hb.TagInterface {
			return bs.FormSelectOption(timezone.Timezone(), timezone.Timezone()).
				AttrIf(selectedTimezone == timezone.Timezone(), "selected", "selected")
		}))

	return selectTimezones
}

func (controller *profileController) prepareData(r *http.Request) (data profileControllerData, errorMessage string) {
	authUser := helpers.GetAuthUser(r)

	if authUser == nil {
		return profileControllerData{}, "User not found"
	}

	countryList, errCountries := config.GeoStore.CountryList(geostore.CountryQueryOptions{
		SortOrder: "asc",
		OrderBy:   geostore.COLUMN_NAME,
	})

	if errCountries != nil {
		config.Logger.Error("Error listing countries", "error", errCountries.Error())
		return profileControllerData{}, "Error listing countries"
	}

	email, firstName, lastName, buinessName, phone, err := controller.untokenizeProfileData(r.Context(), authUser)

	if err != nil {
		config.Logger.Error("Error reading profile data", "error", err.Error())
		return profileControllerData{}, "Error reading profile data"
	}

	data.request = r
	data.authUser = authUser
	data.countryList = countryList

	if r.Method == http.MethodGet {
		data.email = email
		data.firstName = firstName
		data.lastName = lastName
		data.buinessName = buinessName
		data.phone = phone
		data.timezone = authUser.Timezone()
		data.country = authUser.Country()
	}

	if r.Method == http.MethodPost {
		data.email = strings.TrimSpace(utils.Req(r, "email", ""))
		data.firstName = strings.TrimSpace(utils.Req(r, "first_name", ""))
		data.lastName = strings.TrimSpace(utils.Req(r, "last_name", ""))
		data.buinessName = strings.TrimSpace(utils.Req(r, "business_name", ""))
		data.phone = strings.TrimSpace(utils.Req(r, "phone", ""))
		data.timezone = strings.TrimSpace(utils.Req(r, "timezone", ""))
		data.country = strings.TrimSpace(utils.Req(r, "country", ""))
	}

	return data, ""
}

func (controller *profileController) untokenizeProfileData(ctx context.Context, user userstore.UserInterface) (email string, firstName string, lastName string, businessName string, phone string, err error) {
	emailToken := user.Email()
	firstNameToken := user.FirstName()
	lastNameToken := user.LastName()
	businessNameToken := user.BusinessName()
	phoneToken := user.Phone()

	if emailToken != "" {
		email, err = config.VaultStore.TokenRead(ctx, emailToken, config.VaultKey)

		if err != nil {
			config.Logger.Error("Error reading email", "error", err.Error())
			return "", "", "", "", "", err
		}
	}

	if firstNameToken != "" {
		firstName, err = config.VaultStore.TokenRead(ctx, firstNameToken, config.VaultKey)

		if err != nil {
			config.Logger.Error("Error reading first name", "error", err.Error())
			return "", "", "", "", "", err
		}
	}

	if lastNameToken != "" {
		lastName, err = config.VaultStore.TokenRead(ctx, lastNameToken, config.VaultKey)

		if err != nil {
			config.Logger.Error("Error reading last name", "error", err.Error())
			return "", "", "", "", "", err
		}
	}

	if businessNameToken != "" {
		businessName, err = config.VaultStore.TokenRead(ctx, businessNameToken, config.VaultKey)

		if err != nil {
			config.Logger.Error("Error reading business name", "error", err.Error())
			return "", "", "", "", "", err
		}
	}

	if phoneToken != "" {
		phone, err = config.VaultStore.TokenRead(ctx, phoneToken, config.VaultKey)

		if err != nil {
			config.Logger.Error("Error reading phone", "error", err.Error())
			return "", "", "", "", "", err
		}
	}

	return email, firstName, lastName, businessName, phone, nil
}

type profileControllerData struct {
	request     *http.Request
	action      string
	authUser    userstore.UserInterface
	email       string
	firstName   string
	lastName    string
	buinessName string
	phone       string
	// email        string
	country            string
	countryList        []geostore.Country
	timezone           string
	formErrorMessage   string
	formSuccessMessage string
	formRedirectURL    string
}
