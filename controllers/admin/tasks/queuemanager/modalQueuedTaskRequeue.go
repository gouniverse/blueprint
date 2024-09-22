package adminTasks

// import (
// 	"github.com/gouniverse/bs"
// 	"github.com/gouniverse/hb"
// )

// func (controller *queueManagerController) modalQueuedTaskRequeue(parameters string) hb.TagInterface {
// 	modalCloseScript := `document.getElementById('ModalMessage').remove();document.getElementById('ModalBackdrop').remove();`

// 	buttonCancel := hb.NewButton().
// 		Child(hb.NewI().Class("bi bi-chevron-left")).
// 		HTML("Cancel").
// 		Class("btn btn-secondary float-start").
// 		OnClick(modalCloseScript)

// 	buttonRequeue := hb.NewButton().
// 		Child(hb.NewI().Class("bi bi-arrow-clockwise")).
// 		HTML("Requeue").
// 		Class("btn btn-primary float-end")

// 	modal := bs.Modal().
// 		ID("ModalMessage").
// 		Class("fade show modal-lg").
// 		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
// 		Children([]hb.TagInterface{
// 			bs.ModalDialog().Children([]hb.TagInterface{
// 				bs.ModalContent().Children([]hb.TagInterface{
// 					bs.ModalHeader().Children([]hb.TagInterface{
// 						hb.NewHeading5().
// 							HTML("Queued Task Requeue"),
// 						hb.NewButton().
// 							Type("button").
// 							Class("btn-close").
// 							Data("bs-dismiss", "modal").
// 							OnClick(modalCloseScript),
// 					}),

// 					bs.ModalBody().
// 						Child(
// 							hb.NewDiv().
// 								HTML("Queued Task Parameters:").
// 								Style(`font-size:18px;color:black;font-weight:bold;`),
// 						).
// 						Child(
// 							hb.NewTextArea().
// 								Class("form-control").
// 								Style(`height:600px;`).
// 								Name("parameters").
// 								HTML(parameters),
// 						),

// 					bs.ModalFooter().
// 						Style(`display:flex;justify-content:space-between;`).
// 						Child(buttonCancel).
// 						Child(buttonRequeue),
// 				}),
// 			}),
// 		})

// 	backdrop := hb.NewDiv().
// 		ID("ModalBackdrop").
// 		Class("modal-backdrop fade show").
// 		Style("display:block;")

// 	return hb.NewWrap().Children([]hb.TagInterface{
// 		modal,
// 		backdrop,
// 	})
// }
