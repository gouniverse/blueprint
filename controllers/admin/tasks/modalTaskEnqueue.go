package admin

import (
	"net/http"
	"project/config"
	"project/internal/links"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/taskstore"
	"github.com/samber/lo"
)

func (controller *queueManagerController) modalTaskEnqueue(_ *http.Request) *hb.Tag {
	taskList, err := config.TaskStore.TaskList(taskstore.TaskQueryOptions{
		Status:    taskstore.TaskStatusActive,
		OrderBy:   taskstore.COLUMN_TITLE,
		SortOrder: taskstore.DESC,
	})

	if err != nil {
		config.Logger.Error("At adminTasks > modalTaskEnqueue", "error", err.Error())
		return hb.NewSwal(hb.SwalOptions{Title: "Error", Text: "Error listing tasks"})
	}

	groupTasks := bs.FormGroup().
		Class("mb-3").
		Child(bs.FormLabel("Task").
			Style(`font-size:18px;color:black;font-weight:bold;`).
			Child(hb.NewSup().Text("*").Class("text-danger"))).
		Child(bs.FormSelect().
			Name("task_id").
			Child(hb.NewOption().Value("").Text("- Select Task -")).
			Children(lo.Map(taskList, func(task taskstore.Task, _ int) hb.TagInterface {
				return hb.NewOption().
					Value(task.ID).
					Text(task.Title)
			})))

	groupParameters := bs.FormGroup().
		Class("mb-3").
		Child(bs.FormLabel("Task Parameters").
			Style(`font-size:18px;color:black;font-weight:bold;`).
			Child(hb.NewSup().Text("*").Class("text-danger"))).
		Child(bs.FormTextArea().
			Name("task_parameters").
			Class("form-control").
			Style(`height:300px;`).
			Placeholder("Parameters")).
		Child(hb.NewDiv().
			Text("Must be a valid JSON string").
			Class("form-text text-muted"))

	modalCloseScript := `document.getElementById('ModalMessage').remove();document.getElementById('ModalBackdrop').remove();`
	butonModalClose := hb.NewButton().Type("button").
		Class("btn-close").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	buttonCancel := hb.NewButton().
		Child(hb.NewI().Class("bi bi-chevron-left me-2")).
		HTML("Cancel").
		Class("btn btn-secondary float-start").
		OnClick(modalCloseScript)

	buttonEnqueue := hb.NewButton().
		Child(hb.NewI().Class("bi bi-play me-2")).
		HTML("Add to queue").
		Class("btn btn-primary float-end").
		HxInclude(`#ModalTaskEnqueue`).
		HxPost(links.NewAdminLinks().Tasks(map[string]string{
			"action": ActionModalQueuedTaskEnqueueSubmitted,
		})).
		HxTarget("body").
		HxSwap("beforeend")

	modal := bs.Modal().
		ID("ModalTaskEnqueue").
		Class("fade show").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Children([]hb.TagInterface{
			bs.ModalDialog().Children([]hb.TagInterface{
				bs.ModalContent().Children([]hb.TagInterface{
					bs.ModalHeader().Children([]hb.TagInterface{
						hb.NewHeading5().
							Text("New Task Enqueue").
							Style(`padding: 0px; margin: 0px;`),
						butonModalClose,
					}),

					bs.ModalBody().
						Child(groupTasks).
						Child(groupParameters),

					bs.ModalFooter().
						Style(`display:flex;justify-content:space-between;`).
						Child(buttonCancel).
						Child(buttonEnqueue),
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
