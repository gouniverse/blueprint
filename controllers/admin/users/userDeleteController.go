package admin

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/userstore"
	"github.com/gouniverse/utils"
)

type userDeleteController struct{}

var _ router.HTMLControllerInterface = (*userDeleteController)(nil)

type userDeleteControllerData struct {
	userID         string
	user           userstore.UserInterface
	successMessage string
	//errorMessage   string
}

func NewUserDeleteController() *userDeleteController {
	return &userDeleteController{}
}

func (controller userDeleteController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareDataAndValidate(r)

	if errorMessage != "" {
		return hb.Swal(hb.SwalOptions{
			Icon: "error",
			Text: errorMessage,
		}).ToHTML()
	}

	if data.successMessage != "" {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon: "success",
				Text: data.successMessage,
			})).
			Child(hb.Script("setTimeout(() => {window.location.href = window.location.href}, 2000)")).
			ToHTML()
	}

	return controller.
		modal(data).
		ToHTML()
}

func (controller *userDeleteController) modal(data userDeleteControllerData) hb.TagInterface {
	submitUrl := links.NewAdminLinks().UsersUserDelete(map[string]string{
		"user_id": data.userID,
	})

	modalID := "ModalUserDelete"
	modalBackdropClass := "ModalBackdrop"

	formGroupUserId := hb.Input().
		Type(hb.TYPE_HIDDEN).
		Name("user_id").
		Value(data.userID)

	buttonDelete := hb.Button().
		HTML("Delete").
		Class("btn btn-primary float-end").
		HxInclude("#Modal" + modalID).
		HxPost(submitUrl).
		HxSelectOob("#ModalUserDelete").
		HxTarget("body").
		HxSwap("beforeend")

	modalCloseScript := `closeModal` + modalID + `();`

	modalHeading := hb.Heading5().HTML("Delete User").Style(`margin:0px;`)

	modalClose := hb.Button().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	jsCloseFn := `function closeModal` + modalID + `() {document.getElementById('ModalUserDelete').remove();[...document.getElementsByClassName('` + modalBackdropClass + `')].forEach(el => el.remove());}`

	modal := bs.Modal().
		ID(modalID).
		Class("fade show").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Child(hb.Script(jsCloseFn)).
		Child(bs.ModalDialog().
			Child(bs.ModalContent().
				Child(
					bs.ModalHeader().
						Child(modalHeading).
						Child(modalClose)).
				Child(
					bs.ModalBody().
						Child(hb.Paragraph().Text("Are you sure you want to delete this user?").Style(`margin-bottom:20px;color:red;`)).
						Child(hb.Paragraph().Text("This action cannot be undone.")).
						Child(formGroupUserId)).
				Child(bs.ModalFooter().
					Style(`display:flex;justify-content:space-between;`).
					Child(
						hb.Button().HTML("Close").
							Class("btn btn-secondary float-start").
							Data("bs-dismiss", "modal").
							OnClick(modalCloseScript)).
					Child(buttonDelete)),
			))

	backdrop := hb.Div().Class(modalBackdropClass).
		Class("modal-backdrop fade show").
		Style("display:block;z-index:1000;")

	return hb.Wrap().
		Children([]hb.TagInterface{
			modal,
			backdrop,
		})
}

func (controller *userDeleteController) prepareDataAndValidate(r *http.Request) (data userDeleteControllerData, errorMessage string) {
	authUser := helpers.GetAuthUser(r)
	data.userID = utils.Req(r, "user_id", "")

	if authUser == nil {
		return data, "You are not logged in. Please login to continue."
	}

	if data.userID == "" {
		return data, "user id is required"
	}

	user, err := config.UserStore.UserFindByID(r.Context(), data.userID)

	if err != nil {
		config.LogStore.ErrorWithContext("Error. At userDeleteController > prepareDataAndValidate", err.Error())
		return data, "User not found"
	}

	if user == nil {
		return data, "User not found"
	}

	data.user = user

	if r.Method != "POST" {
		return data, ""
	}

	err = config.UserStore.UserSoftDelete(r.Context(), user)

	if err != nil {
		config.LogStore.ErrorWithContext("Error. At userDeleteController > prepareDataAndValidate", err.Error())
		return data, "Deleting user failed. Please contact an administrator."
	}

	data.successMessage = "user deleted successfully."

	return data, ""

}
