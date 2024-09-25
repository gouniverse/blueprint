package admin

import (
	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
)

func (controller *queueManagerController) modalQueuedTaskParameters(parameters string) *hb.Tag {
	modalCloseScript := `document.getElementById('ModalMessage').remove();document.getElementById('ModalBackdrop').remove();`

	title := hb.NewHeading5().
		Text("Queued Task Parameters").
		Style(`margin:0px;padding:0px;`)

	butonModalClose := hb.NewButton().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	groupParameters := bs.FormGroup().
		Child(
			hb.NewDiv().
				HTML("Parameters:").
				Style(`font-size:18px;color:black;font-weight:bold;`),
		).
		Child(
			hb.NewTextArea().
				Class("form-control").
				Style(`height:300px;`).
				Name("parameters").
				HTML(parameters),
		)

	buttonCancel := hb.NewButton().
		Child(hb.NewI().Class("bi bi-chevron-left me-2")).
		HTML("Cancel").
		Class("btn btn-secondary float-start").
		OnClick(modalCloseScript)

	buttonOk := hb.NewButton().
		Child(hb.NewI().Class("bi bi-check me-2")).
		HTML("Ok").
		Class("btn btn-primary float-end").
		OnClick(modalCloseScript)

	modal := bs.Modal().
		ID("ModalMessage").
		Class("fade show modal-lg").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Children([]hb.TagInterface{
			bs.ModalDialog().Children([]hb.TagInterface{
				bs.ModalContent().Children([]hb.TagInterface{
					bs.ModalHeader().Children([]hb.TagInterface{
						title,
						butonModalClose,
					}),

					bs.ModalBody().
						Child(
							groupParameters,
						),

					bs.ModalFooter().
						Style(`display:flex;justify-content:space-between;`).
						Child(buttonCancel).
						Child(buttonOk),
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
