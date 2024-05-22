package user

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
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

type profileController struct{}

type profileControllerData struct {
	authUser    userstore.User
	email       string
	firstName   string
	lastName    string
	buinessName string
	phone       string
	// email        string
	country            string
	timezone           string
	formErrorMessage   string
	formSuccessMessage string
	formRedirectURL    string
}

func NewProfileController() *profileController {
	return &profileController{}
}

func (controller *profileController) Handle(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareData(r)

	if errorMessage != "" {
		return helpers.ToFlashError(w, r, errorMessage, links.NewUserLinks().Home(), 10)
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

	title := hb.NewHeading1().Text("My Profile").Style("margin:30px 0px 30px 0px;")

	paragraph1 := hb.NewParagraph().Text("Please keep your details updated so that we can contact you if you need our help.").Style("margin-bottom:20px;")

	formProfile := controller.userProfileForm(data)

	page := hb.NewSection().
		Child(breadcrumbs).
		Child(
			hb.NewDiv().
				Class("container").
				Child(title).
				Child(paragraph1).
				Child(formProfile).
				Child(hb.NewBR()).
				Child(hb.NewBR()),
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
		return controller.userProfileForm(data).ToHTML()
	}

	if data.lastName == "" {
		data.formErrorMessage = "Last name is required field"
		return controller.userProfileForm(data).ToHTML()
	}

	data.authUser.SetFirstName(data.firstName)
	data.authUser.SetLastName(data.lastName)
	data.authUser.SetBusinessName(data.buinessName)
	data.authUser.SetPhone(data.phone)
	err := config.UserStore.UserUpdate(&data.authUser)

	if err != nil {
		config.LogStore.ErrorWithContext("Error updating user profile", err.Error())

		data.formErrorMessage = "Saving profile failed. Please try again later."
		return controller.userProfileForm(data).ToHTML()
	}

	data.formSuccessMessage = "Profile updated successfully"
	data.formRedirectURL = helpers.ToFlashSuccessURL(data.formSuccessMessage, links.NewUserLinks().Home(), 5)
	return controller.userProfileForm(data).ToHTML()
}

func (controller *profileController) userProfileForm(data profileControllerData) *hb.Tag {
	required := hb.NewSup().
		Text("required").
		Style("margin-left:5px;color:lightcoral;")

	formProfile := hb.NewDiv().ID("FormProfile").Children([]hb.TagInterface{
		bs.Row().Class("g-4").Children([]hb.TagInterface{
			// First name
			bs.Column(6).Children([]hb.TagInterface{
				bs.FormGroup().Children([]hb.TagInterface{
					bs.FormLabel("First name").
						Child(required),
					bs.FormInput().
						Name("first_name").
						Value(data.firstName),
				}),
			}),
			// Last name
			bs.Column(6).Children([]hb.TagInterface{
				bs.FormGroup().Children([]hb.TagInterface{
					bs.FormLabel("Last name").Child(required),
					bs.FormInput().
						Name("last_name").
						Value(data.lastName),
				}),
			}),
			// Buiness / company
			bs.Column(6).Children([]hb.TagInterface{
				bs.FormGroup().Children([]hb.TagInterface{
					bs.FormLabel("Company / buiness name"),
					bs.FormInput().
						Name("business_name").
						Value(data.buinessName),
				}),
			}),
			// Phone
			bs.Column(6).Children([]hb.TagInterface{
				bs.FormGroup().Children([]hb.TagInterface{
					bs.FormLabel("Phone"),
					bs.FormInput().
						Name("phone").
						Value(data.phone),
				}),
			}),
			bs.Column(12).Children([]hb.TagInterface{
				bs.FormGroup().Children([]hb.TagInterface{
					bs.FormLabel("Email").
						Child(required),
					bs.FormInput().
						Name("email").
						Value(data.email).
						Attr("readonly", "readonly").
						Style("background-color:#F8F8F8;"),
				}),
			}),
		}),
		bs.Row().Class("mt-3").Children([]hb.TagInterface{
			bs.Column(12).Class("d-sm-flex justify-content-end").
				Children([]hb.TagInterface{
					// Save Button
					bs.Button().
						Class("btn-primary mb-0").
						Attr("type", "button").
						Text("Save changes").
						HxInclude("#FormProfile").
						HxTarget("#CardUserProfile").
						HxTrigger("click").
						HxSwap("outerHTML").
						HxPost(links.NewUserLinks().Profile()),
				}),
		}),
	})

	// errorMessageJSON, _ := utils.ToJSON(data.formErrorMessage)
	// successMessageJSON, _ := utils.ToJSON(data.formSuccessMessage)
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

func (controller *profileController) prepareData(r *http.Request) (data profileControllerData, errorMessage string) {
	authUser := helpers.GetAuthUser(r)

	if authUser == nil {
		return profileControllerData{}, "User not found"
	}

	if r.Method == http.MethodGet {
		data = profileControllerData{
			authUser:    *authUser,
			email:       authUser.Email(),
			firstName:   authUser.FirstName(),
			lastName:    authUser.LastName(),
			buinessName: authUser.BusinessName(),
			phone:       authUser.Phone(),
			timezone:    authUser.Timezone(),
			country:     authUser.Country(),
		}
	}

	if r.Method == http.MethodPost {
		data = profileControllerData{
			authUser:    *authUser,
			email:       authUser.Email(),
			firstName:   strings.TrimSpace(utils.Req(r, "first_name", "")),
			lastName:    strings.TrimSpace(utils.Req(r, "last_name", "")),
			buinessName: strings.TrimSpace(utils.Req(r, "business_name", "")),
			phone:       strings.TrimSpace(utils.Req(r, "phone", "")),
			timezone:    strings.TrimSpace(utils.Req(r, "timezone", "")),
			country:     strings.TrimSpace(utils.Req(r, "country", "")),
		}
	}

	return data, ""
}
