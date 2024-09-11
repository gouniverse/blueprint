package auth

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/layouts"
	"project/internal/links"
	"project/pkg/userstore"
	"strings"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/geostore"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

// == CONTROLLER ==============================================================

type registerController struct {
	actionOnCountrySelectedTimezoneOptions string
	formFirstName                          string
	formLastName                           string
	formBusinessName                       string
	formEmail                              string
	formPhone                              string
	formCountry                            string
	formTimezone                           string
}

type registerControllerData struct {
	action             string
	authUser           userstore.User
	email              string
	firstName          string
	lastName           string
	buinessName        string
	phone              string
	country            string
	timezone           string
	countryList        []geostore.Country
	formErrorMessage   string
	formSuccessMessage string
	formRedirectURL    string
}

// == CONSTRUCTOR =============================================================

func NewRegisterController() *registerController {
	return &registerController{
		actionOnCountrySelectedTimezoneOptions: "on-country-selected-timezone-options",
		formCountry:                            "country",
		formTimezone:                           "timezone",
		formPhone:                              "phone",
		formEmail:                              "email",
		formFirstName:                          "first_name",
		formLastName:                           "last_name",
		formBusinessName:                       "buiness_name",
	}
}

// == PUBLIC METHODS ==========================================================

func (controller *registerController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareData(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewWebsiteLinks().Home(), 10)
	}

	if data.action == controller.actionOnCountrySelectedTimezoneOptions {
		return controller.selectTimezoneByCountry(data.country, data.timezone).ToHTML()
	}

	if r.Method == http.MethodPost {
		return controller.postUpdate(data)
	}

	return layouts.NewGuestLayout(layouts.Options{
		Title: "Register",
		// CanonicalURL: links.NewWebsiteLinks().Flash(map[string]string{}),
		Content: controller.pageHTML(data),
		ScriptURLs: []string{
			cdn.BootstrapJs_5_3_3(),
			cdn.Htmx_2_0_0(),
			cdn.Sweetalert2_11(),
		},
		StyleURLs: []string{cdn.BootstrapIconsCss_1_11_3()},
		Styles: []string{`.Center > div{padding:0px !important;margin:0px !important;}
		@media (min-width: 576px) {.container.container-xs {max-width: 350px;}}
		body{background:rgba(128,0,128,0.05);}`},
	}).ToHTML()
}

// == PRIVATE METHODS =========================================================

func (controller *registerController) postUpdate(data registerControllerData) string {
	if data.firstName == "" {
		data.formErrorMessage = "First name is required field"
		return controller.formRegister(data).ToHTML()
	}

	if data.lastName == "" {
		data.formErrorMessage = "Last name is required field"
		return controller.formRegister(data).ToHTML()
	}

	if data.country == "" {
		data.formErrorMessage = "Country is required field"
		return controller.formRegister(data).ToHTML()
	}

	if data.timezone == "" {
		data.formErrorMessage = "Timezone is required field"
		return controller.formRegister(data).ToHTML()
	}

	firstNameToken, err := config.VaultStore.TokenCreate(data.firstName, config.VaultKey, 20)

	if err != nil {
		config.LogStore.ErrorWithContext("Error creating first name token", err.Error())
		data.formErrorMessage = "We are very sorry. Saving the details failed. Please try again later."
		return controller.formRegister(data).ToHTML()
	}

	lastNameToken, err := config.VaultStore.TokenCreate(data.lastName, config.VaultKey, 20)

	if err != nil {
		config.LogStore.ErrorWithContext("Error creating last name token", err.Error())
		data.formErrorMessage = "We are very sorry. Saving the details failed. Please try again later."
		return controller.formRegister(data).ToHTML()
	}

	businessNameToken, err := config.VaultStore.TokenCreate(data.buinessName, config.VaultKey, 20)

	if err != nil {
		config.LogStore.ErrorWithContext("Error creating business name token", err.Error())
		data.formErrorMessage = "We are very sorry. Saving the details failed. Please try again later."
		return controller.formRegister(data).ToHTML()
	}

	phoneToken, err := config.VaultStore.TokenCreate(data.phone, config.VaultKey, 20)

	if err != nil {
		config.LogStore.ErrorWithContext("Error creating phone token", err.Error())
		data.formErrorMessage = "We are very sorry. Saving the details failed. Please try again later."
		return controller.formRegister(data).ToHTML()
	}

	data.authUser.SetFirstName(firstNameToken)
	data.authUser.SetLastName(lastNameToken)
	data.authUser.SetBusinessName(businessNameToken)
	data.authUser.SetPhone(phoneToken)
	data.authUser.SetCountry(data.country)
	data.authUser.SetTimezone(data.timezone)

	err = config.UserStore.UserUpdate(&data.authUser)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating user profile", err.Error())

		data.formErrorMessage = "We are very sorry. Saving the details failed. Please try again later."
		return controller.formRegister(data).ToHTML()
	}

	data.formSuccessMessage = "Your registration completed successfully. You can now continue browsing the website."
	data.formRedirectURL = helpers.ToFlashSuccessURL(data.formSuccessMessage, links.NewUserLinks().Home(), 5)
	return controller.formRegister(data).ToHTML()
}

func (controller *registerController) pageHTML(data registerControllerData) *hb.Tag {
	form := controller.formRegister(data)
	return hb.NewDiv().
		Class(`container container-xs text-center`).
		Child(hb.NewBR()).
		Child(hb.NewBR()).
		Child(hb.NewHTML(layouts.LogoHTML())).
		Child(hb.NewBR()).
		Child(hb.NewBR()).
		Child(hb.NewHeading1().Text("Complete registration").Style(`font-size:24px;`)).
		Child(hb.NewBR()).
		Child(form).
		Child(hb.NewBR()).
		Child(hb.NewBR())
}

func (controller *registerController) formRegister(data registerControllerData) *hb.Tag {
	required := hb.NewSup().
		Text("required").
		Style("margin-left:5px;color:lightcoral;")

	buttonSave := bs.Button().
		Class("btn-primary mb-0 w-100").
		Attr("type", "button").
		Child(hb.NewI().Class("bi bi-check-circle me-2")).
		Text("Save changes").
		HxInclude("#FormRegister").
		HxTarget("#CardRegister").
		HxTrigger("click").
		HxSwap("outerHTML").
		HxPost(links.NewAuthLinks().Register(map[string]string{}))

	firstNameGroup := hb.NewDiv().
		Class("form-group").
		Children([]hb.TagInterface{
			bs.FormLabel("First name").
				Child(required),
			bs.FormInput().
				Name(controller.formFirstName).
				Value(data.firstName),
		})

	lastNameGroup := hb.NewDiv().
		Class("form-group").
		Children([]hb.TagInterface{
			bs.FormLabel("Last name").
				Child(required),
			bs.FormInput().
				Name(controller.formLastName).
				Value(data.lastName),
		})

	// businessNameGroup := bs.FormGroup().Children([]hb.TagInterface{
	// 	bs.FormLabel("Company / buiness name"),
	// 	bs.FormInput().
	// 		Name("business_name").
	// 		Value(data.buinessName),
	// })

	// phoneGroup := bs.FormGroup().Children([]hb.TagInterface{
	// 	bs.FormLabel("Phone"),
	// 	bs.FormInput().
	// 		Name("phone").
	// 		Value(data.phone),
	// })

	emailGroup := hb.NewDiv().
		Class("form-group").
		Children([]hb.TagInterface{
			bs.FormLabel("Email").
				Child(required),
			bs.FormInput().
				Name("email").
				Value(data.email).
				Attr("readonly", "readonly").
				Style("background-color:#F8F8F8;"),
		})

	selectCountries := bs.FormSelect().
		ID("SelectCountries").
		Name(controller.formCountry).
		Child(bs.FormSelectOption("", "")).
		Children(lo.Map(data.countryList, func(country geostore.Country, _ int) hb.TagInterface {
			return bs.FormSelectOption(country.IsoCode2(), country.Name()).
				AttrIf(data.country == country.IsoCode2(), "selected", "selected")
		})).
		Hx("post", links.NewAuthLinks().Register(map[string]string{
			"action": "on-country-selected-timezone-options",
		})).
		Hx("target", "#SelectTimezones").
		Hx("swap", "outerHTML")

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

	formProfile := hb.NewDiv().
		ID("FormRegister").
		Child(
			bs.Row().
				Class("g-4").
				Children([]hb.TagInterface{
					bs.Column(12).
						Class("mt-3").
						Child(emailGroup),
					bs.Column(6).
						Class("mt-2").
						Child(firstNameGroup),
					bs.Column(6).
						Class("mt-2").
						Child(lastNameGroup),
					// bs.Column(6).
					// 	Child(businessNameGroup),
					// bs.Column(6).
					// 	Child(phoneGroup),
					bs.Column(6).
						Class("mt-2").
						Child(countryGroup),
					bs.Column(6).
						Class("mt-2").
						Child(timezoneGroup),
				}),
		).
		Child(
			bs.Row().Class("mt-3").Children([]hb.TagInterface{
				bs.Column(12).Class("d-sm-flex justify-content-end").
					Children([]hb.TagInterface{
						buttonSave,
					}),
			}),
		)

	return hb.NewDiv().ID("CardRegister").
		Class("card bg-white border rounded-3").
		Style("text-align:left;").
		Children([]hb.TagInterface{
			hb.NewDiv().Class("card-header  bg-transparent").Children([]hb.TagInterface{
				hb.NewHeading3().
					Text("Your Details").
					Style("text-align:left;font-size:12px;color:#333;margin:0px;"),
			}),
			hb.NewDiv().Class("card-body").Children([]hb.TagInterface{
				formProfile,
			}),
		}).
		ChildIf(data.formErrorMessage != "", hb.NewSwal(hb.SwalOptions{
			Icon:              "error",
			Title:             "Oops...",
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

func (controller *registerController) untokenizeData(user userstore.User) (email string, firstName string, lastName string, businessName string, phone string, err error) {
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

func (controller *registerController) prepareData(r *http.Request) (data registerControllerData, errorMessage string) {
	action := utils.Req(r, "action", "")
	authUser := helpers.GetAuthUser(r)

	if authUser == nil {
		return registerControllerData{}, "User not found"
	}

	countries, errCountries := config.GeoStore.CountryList(geostore.CountryQueryOptions{
		SortOrder: "asc",
		OrderBy:   geostore.COLUMN_NAME,
	})

	if errCountries != nil {
		config.LogStore.ErrorWithContext("Error listing countries", errCountries.Error())
		return registerControllerData{}, "Error listing countries"
	}

	email, firstName, lastName, businessName, phone, err := controller.untokenizeData(*authUser)

	if r.Method == http.MethodGet {
		if err != nil {
			config.LogStore.ErrorWithContext("Error reading email", err.Error())
			return registerControllerData{}, "Error reading email"
		}

		data = registerControllerData{
			action:      action,
			authUser:    *authUser,
			email:       email,
			firstName:   firstName,
			lastName:    lastName,
			buinessName: businessName,
			phone:       phone,
			timezone:    authUser.Timezone(),
			country:     authUser.Country(),
			countryList: countries,
		}
	}

	if r.Method == http.MethodPost {
		data = registerControllerData{
			action:      action,
			authUser:    *authUser,
			email:       email,
			firstName:   strings.TrimSpace(utils.Req(r, controller.formFirstName, "")),
			lastName:    strings.TrimSpace(utils.Req(r, controller.formLastName, "")),
			buinessName: strings.TrimSpace(utils.Req(r, controller.formBusinessName, "")),
			phone:       strings.TrimSpace(utils.Req(r, controller.formPhone, "")),
			timezone:    strings.TrimSpace(utils.Req(r, controller.formTimezone, "")),
			country:     strings.TrimSpace(utils.Req(r, controller.formCountry, "")),
			countryList: countries,
		}
	}

	return data, ""
}

func (controller *registerController) selectTimezoneByCountry(country string, selectedTimezone string) *hb.Tag {
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
		Name("timezone").
		Child(bs.FormSelectOption("", "")).
		Children(lo.Map(timezones, func(timezone geostore.Timezone, _ int) hb.TagInterface {
			return bs.FormSelectOption(timezone.Timezone(), timezone.Timezone()).
				AttrIf(selectedTimezone == timezone.Timezone(), "selected", "selected")
		}))

	return selectTimezones
}