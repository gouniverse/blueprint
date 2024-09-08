package admin

import (
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"

	"github.com/gouniverse/blogstore"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
)

type postDeleteController struct{}

var _ router.ControllerInterface = (*postDeleteController)(nil)

type postDeleteControllerData struct {
	postID         string
	post           *blogstore.Post
	successMessage string
	//errorMessage   string
}

func NewPostDeleteController() *postDeleteController {
	return &postDeleteController{}
}

func (controller postDeleteController) Handler(w http.ResponseWriter, r *http.Request) string {
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

func (controller *postDeleteController) modal(data postDeleteControllerData) hb.TagInterface {
	submitUrl := links.NewAdminLinks().BlogPostDelete(map[string]string{
		"post_id": data.postID,
	})

	modalID := "ModalPostDelete"
	modalBackdropClass := "ModalBackdrop"

	formGroupPostId := hb.NewInput().
		Type(hb.TYPE_HIDDEN).
		Name("post_id").
		Value(data.postID)

	buttonDelete := hb.NewButton().
		HTML("Delete").
		Class("btn btn-primary float-end").
		HxInclude("#Modal" + modalID).
		HxPost(submitUrl).
		HxSelectOob("#ModalPostDelete").
		HxTarget("body").
		HxSwap("beforeend")

	modalCloseScript := `closeModal` + modalID + `();`

	modalHeading := hb.NewHeading5().HTML("Delete Post").Style(`margin:0px;`)

	modalClose := hb.NewButton().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	jsCloseFn := `function closeModal` + modalID + `() {document.getElementById('ModalPostDelete').remove();[...document.getElementsByClassName('` + modalBackdropClass + `')].forEach(el => el.remove());}`

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
						Child(hb.NewParagraph().Text("Are you sure you want to delete this post?").Style(`margin-bottom:20px;color:red;`)).
						Child(hb.NewParagraph().Text("This action cannot be undone.")).
						Child(formGroupPostId)).
				Child(bs.ModalFooter().
					Style(`display:flex;justify-content:space-between;`).
					Child(
						hb.NewButton().HTML("Close").
							Class("btn btn-secondary float-start").
							Data("bs-dismiss", "modal").
							OnClick(modalCloseScript)).
					Child(buttonDelete)),
			))

	backdrop := hb.NewDiv().Class(modalBackdropClass).
		Class("modal-backdrop fade show").
		Style("display:block;z-index:1000;")

	return hb.NewWrap().
		Children([]hb.TagInterface{
			modal,
			backdrop,
		})
}

func (controller *postDeleteController) prepareDataAndValidate(r *http.Request) (data postDeleteControllerData, errorMessage string) {
	authUser := helpers.GetAuthUser(r)
	data.postID = utils.Req(r, "post_id", "")

	if authUser == nil {
		return data, "You are not logged in. Please login to continue."
	}

	if data.postID == "" {
		return data, "post id is required"
	}

	post, err := config.BlogStore.PostFindByID(data.postID)

	if err != nil {
		config.LogStore.ErrorWithContext("Error. At postDeleteController > prepareDataAndValidate", err.Error())
		return data, "Post not found"
	}

	if post == nil {
		return data, "Post not found"
	}

	data.post = post

	if r.Method != "POST" {
		return data, ""
	}

	err = config.BlogStore.PostTrash(post)

	if err != nil {
		config.LogStore.ErrorWithContext("Error. At postDeleteController > prepareDataAndValidate", err.Error())
		return data, "Deleting post failed. Please contact an administrator."
	}

	data.successMessage = "post deleted successfully."

	return data, ""

}
