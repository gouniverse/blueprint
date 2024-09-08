package admin

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
	"strings"

	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
)

type postCreateController struct{}

var _ router.ControllerInterface = (*postCreateController)(nil)

type postCreateControllerData struct {
	title          string
	successMessage string
	//errorMessage   string
}

func NewPostCreateController() *postCreateController {
	return &postCreateController{}
}

func (controller postCreateController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareDataAndValidate(r)

	if errorMessage != "" {
		return hb.NewSwal(hb.SwalOptions{
			Icon: "error",
			Text: errorMessage,
		}).ToHTML()
	}

	if data.successMessage != "" {
		return hb.NewWrap().
			Child(hb.NewSwal(hb.SwalOptions{
				Icon: "success",
				Text: data.successMessage,
			})).
			Child(hb.NewScript("setTimeout(() => {window.location.href = window.location.href}, 2000)")).
			ToHTML()
	}

	return controller.
		modal(data).
		ToHTML()
}

func (controller *postCreateController) modal(data postCreateControllerData) hb.TagInterface {
	submitUrl := links.NewAdminLinks().BlogPostCreate(map[string]string{})

	formGroupTitle := bs.FormGroup().
		Class("mb-3").
		Child(bs.FormLabel("Title")).
		Child(bs.FormInput().Name("post_title").Value(data.title))

	modalID := "ModalPostCreate"
	modalBackdropClass := "ModalBackdrop"

	modalCloseScript := `closeModal` + modalID + `();`

	modalHeading := hb.NewHeading5().HTML("New Post Create").Style(`margin:0px;`)

	modalClose := hb.NewButton().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	jsCloseFn := `function closeModal` + modalID + `() {document.getElementById('ModalPostCreate').remove();[...document.getElementsByClassName('` + modalBackdropClass + `')].forEach(el => el.remove());}`

	buttonSend := hb.NewButton().
		Child(hb.NewI().Class("bi bi-check me-2")).
		HTML("Create & Edit").
		Class("btn btn-primary float-end").
		HxInclude("#" + modalID).
		HxPost(submitUrl).
		HxSelectOob("#ModalPostCreate").
		HxTarget("body").
		HxSwap("beforeend")

	buttonCancel := hb.NewButton().
		Child(hb.NewI().Class("bi bi-chevron-left me-2")).
		HTML("Close").
		Class("btn btn-secondary float-start").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	modal := bs.Modal().
		ID(modalID).
		Class("fade show").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Child(hb.NewScript(jsCloseFn)).
		Child(bs.ModalDialog().
			Child(bs.ModalContent().
				Child(
					bs.ModalHeader().
						Child(modalHeading).
						Child(modalClose)).
				Child(
					bs.ModalBody().
						Child(formGroupTitle)).
				Child(bs.ModalFooter().
					Style(`display:flex;justify-content:space-between;`).
					Child(buttonCancel).
					Child(buttonSend)),
			))

	backdrop := hb.NewDiv().Class(modalBackdropClass).
		Class("modal-backdrop fade show").
		Style("display:block;z-index:1000;")

	return hb.NewWrap().Children([]hb.TagInterface{
		modal,
		backdrop,
	})
}

func (controller *postCreateController) prepareDataAndValidate(r *http.Request) (data postCreateControllerData, errorMessage string) {
	authUser := helpers.GetAuthUser(r)

	if authUser == nil {
		return data, "You are not logged in. Please login to continue."
	}

	data.title = strings.TrimSpace(utils.Req(r, "post_title", ""))

	if r.Method != "POST" {
		return data, ""
	}

	if data.title == "" {
		return data, "post title is required"
	}

	post := blogstore.NewPost()
	post.SetTitle(data.title)

	err := config.BlogStore.PostCreate(post)

	if err != nil {
		config.LogStore.ErrorWithContext("Error. At postCreateController > prepareDataAndValidate", err.Error())
		return data, "Creating post failed. Please contact an administrator."
	}

	data.successMessage = "post created successfully."

	return data, ""

}
