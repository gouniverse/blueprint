package user

import (
	"net/http"
	"project/config"
	"project/controllers/user/partials"
	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"
	"project/pkg/userstore"
	"strings"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/sb"
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

type profileControllerData struct {
	request     *http.Request
	action      string
	authUser    userstore.User
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
		return helpers.ToFlashError(w, r, errorMessage, links.NewUserLinks().Home(), 10)
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
			URL:  links.NewUserLinks().Profile(),
		},
	})

	title := hb.NewHeading1().
		Text("My Account").
		Style("margin:30px 0px 30px 0px;")

	paragraph1 := hb.NewParagraph().
		Text("Please keep your details updated so that we can contact you if you need our help.").
		Style("margin-bottom:20px;")

	formProfile := controller.formProfile(data)

	page := hb.NewSection().
		Child(breadcrumbs).
		Child(hb.NewHR()).
		Child(partials.UserQuickLinks(data.request)).
		Child(hb.NewHR()).
		Child(
			hb.NewDiv().
				Class("container").
				Child(title).
				Child(paragraph1).
				Child(formProfile).
				Child(hb.NewBR()).
				Child(hb.NewBR()),
			// This feature is not ready, so keep it out until its fleshed out
			// Child(controller.userSubscriptions(r, data)),
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
	err := config.VaultStore.TokenUpdate(data.authUser.FirstName(), data.firstName, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error saving first name", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	// Last name
	err = config.VaultStore.TokenUpdate(data.authUser.LastName(), data.lastName, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error saving last name", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	// Business name
	err = config.VaultStore.TokenUpdate(data.authUser.BusinessName(), data.buinessName, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error saving business name", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	// Phone
	err = config.VaultStore.TokenUpdate(data.authUser.Phone(), data.phone, config.VaultKey)

	if err != nil {
		config.LogStore.ErrorWithContext("Error saving phone", err.Error())
		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	data.authUser.SetCountry(data.country)
	data.authUser.SetTimezone(data.timezone)

	err = config.UserStore.UserUpdate(&data.authUser)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating user profile", err.Error())

		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.formProfile(data).ToHTML()
	}

	data.formSuccessMessage = "Profile updated successfully"
	data.formRedirectURL = helpers.ToFlashSuccessURL(data.formSuccessMessage, links.NewUserLinks().Home(), 5)
	return controller.formProfile(data).ToHTML()
}

func (controller *profileController) formProfile(data profileControllerData) *hb.Tag {
	required := hb.NewSup().
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

	countryGroup := hb.NewDiv().
		Class("form-group").
		Children([]hb.TagInterface{
			bs.FormLabel("Country").
				Child(required),
			selectCountries,
		})

	timezoneGroup := hb.NewDiv().
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
		HxPost(links.NewUserLinks().Profile())

	formProfile := hb.NewDiv().ID("FormProfile").Children([]hb.TagInterface{
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

	return hb.NewDiv().ID("CardUserProfile").
		Class("card bg-transparent border rounded-3").
		Style("text-align:left;").
		Children([]hb.TagInterface{
			hb.NewDiv().Class("card-header  bg-transparent").Children([]hb.TagInterface{
				hb.NewHeading3().
					Text("Your Details").
					Style("text-align:left;font-size:23px;color:#333;"),
			}),
			hb.NewDiv().Class("card-body").Children([]hb.TagInterface{
				formProfile,
			}),
		}).
		ChildIf(data.formErrorMessage != "", hb.NewSwal(hb.SwalOptions{
			Icon:              "error",
			Title:             "Error",
			Text:              data.formErrorMessage,
			ShowCancelButton:  false,
			ConfirmButtonText: "OK",
		})).
		ChildIf(data.formSuccessMessage != "", hb.NewSwal(hb.SwalOptions{
			Icon:              "success",
			Title:             "Saved",
			Text:              data.formSuccessMessage,
			ShowCancelButton:  false,
			ConfirmButtonText: "OK",
			ConfirmCallback:   "window.location.href = window.location.href",
		})).
		ChildIf(data.formRedirectURL != "", hb.NewScript(`window.location.href = '`+data.formRedirectURL+`'`))

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
		config.LogStore.ErrorWithContext("Error listing timezones", errZones.Error())
		return hb.NewHTML("Error listing timezones")
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
		config.LogStore.ErrorWithContext("Error listing countries", errCountries.Error())
		return profileControllerData{}, "Error listing countries"
	}

	email, firstName, lastName, buinessName, phone, err := controller.untokenizeProfileData(*authUser)

	if err != nil {
		config.LogStore.ErrorWithContext("Error reading profile data", err.Error())
		return profileControllerData{}, "Error reading profile data"
	}

	data.request = r
	data.authUser = *authUser
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

func (controller *profileController) untokenizeProfileData(user userstore.User) (email string, firstName string, lastName string, businessName string, phone string, err error) {
	emailToken := user.Email()
	firstNameToken := user.FirstName()
	lastNameToken := user.LastName()
	businessNameToken := user.BusinessName()
	phoneToken := user.Phone()

	if emailToken != "" {
		email, err = config.VaultStore.TokenRead(emailToken, config.VaultKey)

		if err != nil {
			config.LogStore.ErrorWithContext("Error reading email", err.Error())
			return "", "", "", "", "", err
		}
	}

	if firstNameToken != "" {
		firstName, err = config.VaultStore.TokenRead(firstNameToken, config.VaultKey)

		if err != nil {
			config.LogStore.ErrorWithContext("Error reading first name", err.Error())
			return "", "", "", "", "", err
		}
	}

	if lastNameToken != "" {
		lastName, err = config.VaultStore.TokenRead(lastNameToken, config.VaultKey)

		if err != nil {
			config.LogStore.ErrorWithContext("Error reading last name", err.Error())
			return "", "", "", "", "", err
		}
	}

	if businessNameToken != "" {
		businessName, err = config.VaultStore.TokenRead(businessNameToken, config.VaultKey)

		if err != nil {
			config.LogStore.ErrorWithContext("Error reading business name", err.Error())
			return "", "", "", "", "", err
		}
	}

	if phoneToken != "" {
		phone, err = config.VaultStore.TokenRead(phoneToken, config.VaultKey)

		if err != nil {
			config.LogStore.ErrorWithContext("Error reading phone", err.Error())
			return "", "", "", "", "", err
		}
	}

	return email, firstName, lastName, businessName, phone, nil
}
