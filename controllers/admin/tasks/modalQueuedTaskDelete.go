package admin

import (
	"net/http"
	"project/internal/links"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
)

func (controller *queueManagerController) modalQueuedTaskDelete(_ *http.Request, queueID string) *hb.Tag {
	if queueID == "" {
		return hb.NewSwal(hb.SwalOptions{Icon: "error", Title: "Error", Text: "queue id is required"})
	}

	modalCloseScript := `document.getElementById('ModalQueuedTaskDelete').remove();document.getElementById('ModalBackdrop').remove();`

	buttonModalClose := hb.NewButton().
		Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	title := hb.NewHeading5().
		Text("Delete Queued Task").
		Style(`margin:0px;padding:0px;`)

	buttonCancel := hb.NewButton().
		Child(hb.NewI().Class("bi bi-chevron-left me-2")).
		HTML("Cancel").
		Class("btn btn-secondary float-start").
		OnClick(modalCloseScript)

	buttonDelete := hb.NewButton().
		Child(hb.NewI().Class("bi bi-trash me-2")).
		HTML("Delete").
		Class("btn btn-danger float-end").
		HxInclude(`#ModalQueuedTaskDelete`).
		HxPost(links.NewAdminLinks().Tasks(map[string]string{
			"action": ActionModalQueuedTaskDeleteSubmitted,
		})).
		HxTarget("body").
		HxSwap("beforeend")

	groupDetails := hb.NewDiv().
		Class("text-danger").
		Text("Are you sure you want to delete this queued task?").
		Child(hb.NewBR()).
		Text("This action cannot be undone.")

	inputTaskID := hb.NewInput().Type(hb.TYPE_HIDDEN).Name("queue_id").Value(queueID)

	modal := bs.Modal().
		ID("ModalQueuedTaskDelete").
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
						Child(groupDetails).
						Child(inputTaskID),

					bs.ModalFooter().
						Style(`display:flex;justify-content:space-between;`).
						Child(buttonCancel).
						Child(buttonDelete),
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
