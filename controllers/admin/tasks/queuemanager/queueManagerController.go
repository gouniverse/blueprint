package adminTasks

// import (
// 	"net/http"
// 	"project/config"
// 	"project/controllers/admin/partials"
// 	"project/internal/helpers"
// 	"project/internal/links"
// 	"project/models"
// 	"strconv"
// 	"strings"

// 	"github.com/golang-module/carbon/v2"
// 	"github.com/gouniverse/bs"
// 	"github.com/gouniverse/cdn"
// 	"github.com/gouniverse/hb"
// 	"github.com/gouniverse/taskstore"
// 	"github.com/gouniverse/utils"
// 	"github.com/mingrammer/cfmt"
// 	"github.com/samber/lo"
// )

// type queueManagerController struct {
// }

// type queueManagerControllerData struct {
// 	action    string
// 	page      string
// 	sortOrder string
// 	sortBy    string
// 	status    string
// 	search    string
// 	dateFrom  string
// 	dateTo    string
// 	queueID   string
// 	taskID    string
// }

// func NewQueueManagerController() *queueManagerController {
// 	return &queueManagerController{}
// }

// func (controller *queueManagerController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
// 	data, errMessage := controller.prepareData(r)

// 	if errMessage != "" {
// 		return helpers.ToFlashError(w, r, errMessage, links.NewAdminLinks().Home(map[string]string{}), 10)
// 	}

// 	if data.action == "queue-task-delete" {
// 		return controller.queuedTaskDelete(w, r, data)
// 	}

// 	if data.action == "queue-task-details" {
// 		return controller.queuedTaskDetails(w, r, data)
// 	}

// 	if data.action == "queue-task-parameters" {
// 		return controller.queuedTaskParameters(w, r, data)
// 	}

// 	if data.action == "queue-task-requeue" {
// 		return controller.queueTaskRequeue(w, r, data)
// 	}

// 	if data.action == "queue-task-restart" {
// 		return controller.queueTaskRestart(w, r, data)
// 	}

// 	title := hb.NewHeading1().Text("Queue Manager")

// 	breadcrumbs := bs.Breadcrumbs([]bs.Breadcrumb{
// 		{
// 			Name: "Home",
// 			URL:  links.NewAdminLinks().Home(map[string]string{}),
// 		},
// 		{
// 			Name: "Tasks",
// 			URL:  links.NewAdminLinks().Tasks(map[string]string{}),
// 		},
// 		{
// 			Name: "Task Manager",
// 			URL:  links.NewAdminLinks().Tasks(map[string]string{}),
// 		},
// 	})

// 	webpage := hb.NewWrap().Children([]hb.TagInterface{
// 		hb.NewSection().Children([]hb.TagInterface{
// 			hb.NewDiv().Class("container").
// 				Style("padding:40px;").
// 				Children([]hb.TagInterface{
// 					title,
// 					hb.NewDiv().Style("margin-top:20px;margin-bottom:20px;").Children([]hb.TagInterface{
// 						breadcrumbs,
// 					}),
// 					controller.table(data),
// 				}),
// 		}),
// 	})

// 	return partials.AdminDashboard(r, partials.AdminDashboardOptions{
// 		Title:   "Task Manager",
// 		Content: webpage.ToHTML(),
// 		ScriptURLs: []string{
// 			cdn.Htmx_1_9_6(),
// 		},
// 	}).ToHTML()
// }

// func (controller *queueManagerController) queuedTaskDelete(w http.ResponseWriter, r *http.Request, data queueManagerControllerData) string {
// 	if data.queueID == "" {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue id is required").ToHTML()
// 	}

// 	err := config.TaskStore.QueueSoftDeleteByID(data.queueID)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > queueTaskDelete", err.Error())
// 		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be deleted").ToHTML()
// 	}

// 	return controller.table(data).ToHTML()
// }

// func (controller *queueManagerController) queuedTaskDetails(w http.ResponseWriter, r *http.Request, data queueManagerControllerData) string {
// 	if data.queueID == "" {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue id is required").ToHTML()
// 	}

// 	queue, err := config.TaskStore.QueueFindByID(data.queueID)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > queueTaskDelete", err.Error())
// 		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be deleted").ToHTML()
// 	}

// 	if queue == nil {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue not found").ToHTML()
// 	}

// 	return controller.modalQueuedTaskDetails(queue.Details).ToHTML()
// }

// func (controller *queueManagerController) queuedTaskParameters(w http.ResponseWriter, r *http.Request, data queueManagerControllerData) string {
// 	if data.queueID == "" {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue id is required").ToHTML()
// 	}

// 	queue, err := config.TaskStore.QueueFindByID(data.queueID)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > queueTaskDelete", err.Error())
// 		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be deleted").ToHTML()
// 	}

// 	if queue == nil {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue not found").ToHTML()
// 	}

// 	return controller.modalQueuedTaskParameters(queue.Parameters).ToHTML()
// }

// func (controller *queueManagerController) queueTaskRequeue(w http.ResponseWriter, r *http.Request, data queueManagerControllerData) string {
// 	queue, err := config.TaskStore.QueueFindByID(data.queueID)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > queueTaskDelete", err.Error())
// 		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be deleted").ToHTML()
// 	}

// 	if queue == nil {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue not found").ToHTML()
// 	}

// 	return controller.modalQueuedTaskRequeue(queue.Parameters).ToHTML()
// }

// func (controller *queueManagerController) queueTaskRestart(w http.ResponseWriter, r *http.Request, data queueManagerControllerData) string {
// 	if data.queueID == "" {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue id is required").ToHTML()
// 	}

// 	queue, err := config.TaskStore.QueueFindByID(data.queueID)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > queueTaskDelete", err.Error())
// 		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be deleted").ToHTML()
// 	}

// 	if queue == nil {
// 		return hb.NewDiv().Class("alert alert-danger").Text("queue not found").ToHTML()
// 	}

// 	queue.Status = taskstore.QueueStatusQueued

// 	err = config.TaskStore.QueueUpdate(queue)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > queueTaskDelete", err.Error())
// 		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be queued").ToHTML()
// 	}

// 	return controller.table(data).ToHTML()
// }

// func (controller *queueManagerController) table(data queueManagerControllerData) hb.TagInterface {
// 	allTasks, err := config.TaskStore.TaskList(taskstore.TaskQueryOptions{})

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > table", err.Error())
// 		return hb.NewDiv().Text("error retrieving tasks")
// 	}

// 	pageInt, _ := strconv.Atoi(data.page)
// 	perPage := 20

// 	query := taskstore.QueueQueryOptions{
// 		Offset: int64(pageInt * perPage),
// 		Limit:  perPage,
// 		Status: data.status,
// 		// CreatedAtGreaterThan: controller.dateFrom,
// 		// CreatedAtLessThan:    controller.dateTo,
// 		// Search:               controller.search,
// 		SortOrder: data.sortOrder,
// 		SortBy:    data.sortBy,
// 	}

// 	queuedTaskList, err := config.TaskStore.
// 		// EnableDebug(true).
// 		QueueList(query)

// 	if err != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > table", err.Error())
// 		return hb.NewDiv().Text("error retrieving queued tasks")
// 	}

// 	queudTaskCount, errCount := config.TaskStore.QueueCount(query)

// 	if errCount != nil {
// 		config.Cms.LogStore.ErrorWithContext("At adminTasks > table", errCount.Error())
// 		return hb.NewDiv().Text("error retrieving tasks")
// 	}

// 	table := hb.NewTable().
// 		Class("table table-striped table-hover table-bordered").
// 		Children([]hb.TagInterface{
// 			hb.NewThead().Children([]hb.TagInterface{
// 				hb.NewTR().Children([]hb.TagInterface{
// 					hb.NewTH().
// 						Child(controller.sortableColumnLabel(data, "Name, Alias, ID", "id")).
// 						Style(""),
// 					hb.NewTH().
// 						Child(controller.sortableColumnLabel(data, "Start Time", "started_at")).
// 						Style("width: 1px;"),
// 					hb.NewTH().
// 						Child(controller.sortableColumnLabel(data, "End Time", "completed_at")).
// 						Style("width: 1px;"),
// 					hb.NewTH().
// 						HTML("Elapsed Time").
// 						Style("width: 1px;"),
// 					hb.NewTH().
// 						Child(controller.sortableColumnLabel(data, "Status", "status")).
// 						Style("width: 100px;"),
// 					hb.NewTH().
// 						Child(controller.sortableColumnLabel(data, "Action", "role")).
// 						Style("width: 160px;"),
// 				}),
// 			}),
// 			hb.NewTbody().Children(lo.Map(queuedTaskList, func(queuedTask taskstore.Queue, _ int) hb.TagInterface {
// 				task, taskExists := lo.Find(allTasks, func(t taskstore.Task) bool {
// 					return t.ID == queuedTask.TaskID
// 				})

// 				taskName := lo.IfF(taskExists, func() string { return task.Title }).Else("Unknown")

// 				buttonDelete := hb.NewButton().
// 					Class("btn btn-sm btn-danger").
// 					Style("margin-bottom: 2px; margin-left:2px; margin-right:2px;").
// 					HTML("Delete").
// 					HxPost(links.NewAdminLinks().Tasks(map[string]string{
// 						"action":   "queue-task-delete",
// 						"queue_id": queuedTask.ID,
// 						"page":     data.page,
// 						"by":       data.sortBy,
// 						"sort":     data.sortOrder,
// 					})).
// 					HxTarget("#QueuedTasksListTable").
// 					HxSwap("outerHTML")

// 				buttonParameters := hb.NewButton().
// 					Class("btn btn-sm btn-info").
// 					Style("margin-bottom: 2px; margin-left:2px; margin-right:2px;").
// 					HTML("Parameters").
// 					HxPost(links.NewAdminLinks().Tasks(map[string]string{
// 						"action":   "queue-task-parameters",
// 						"queue_id": queuedTask.ID,
// 						"page":     data.page,
// 						"by":       data.sortBy,
// 						"sort":     data.sortOrder,
// 					})).
// 					// HxTarget("#QueuedTasksListTable").
// 					HxTarget("body").
// 					HxSwap("beforeend")

// 				buttonDetails := hb.NewButton().
// 					Class("btn btn-sm btn-info").
// 					Style("margin-bottom: 2px; margin-left:2px; margin-right:2px;").
// 					HTML("Details").
// 					HxPost(links.NewAdminLinks().Tasks(map[string]string{
// 						"action":   "queue-task-details",
// 						"queue_id": queuedTask.ID,
// 						"page":     data.page,
// 						"by":       data.sortBy,
// 						"sort":     data.sortOrder,
// 					})).
// 					// HxTarget("#QueuedTasksListTable").
// 					// HxSelectOob("#ModalMessage").
// 					HxTarget("body").
// 					HxSwap("beforeend")

// 				buttonRequeue := hb.NewButton().
// 					Class("btn btn-sm btn-info").
// 					Style("margin-bottom: 2px; margin-left:2px; margin-right:2px;").
// 					HTML("Requeue").
// 					HxPost(links.NewAdminLinks().Tasks(map[string]string{
// 						"action":   "queue-task-requeue",
// 						"queue_id": queuedTask.ID,
// 						"page":     data.page,
// 						"by":       data.sortBy,
// 						"sort":     data.sortOrder,
// 					})).
// 					HxTarget("body").
// 					HxSwap("beforeend")

// 				buttonRestart := hb.NewButton().
// 					Class("btn btn-sm btn-info").
// 					Style("margin-bottom: 2px; margin-left:2px; margin-right:2px;").
// 					HTML("Restart").
// 					HxPost(links.NewAdminLinks().Tasks(map[string]string{
// 						"action":   "queue-task-restart",
// 						"queue_id": queuedTask.ID,
// 						"page":     data.page,
// 						"by":       data.sortBy,
// 						"sort":     data.sortOrder,
// 					})).
// 					HxTarget("#QueuedTasksListTable").
// 					HxSwap("outerHTML")

// 				// linkTask := hb.NewHyperlink().
// 				// 	HTML(queuedTask.ID).
// 				// 	Href(links.NewAdminLinks().Tasks(map[string]string{
// 				// 		"task_id": queuedTask.ID,
// 				// 	}))

// 				startedAtDate := lo.IfF(queuedTask.StartedAt != nil, func() string {
// 					return carbon.CreateFromStdTime(*queuedTask.StartedAt).Format("d M Y")
// 				}).Else("-")
// 				startedAtTime := lo.IfF(queuedTask.StartedAt != nil, func() string {
// 					return carbon.CreateFromStdTime(*queuedTask.StartedAt).ToTimeString()
// 				}).Else("-")
// 				completeddAtDate := lo.IfF(queuedTask.CompletedAt != nil, func() string {
// 					return carbon.CreateFromStdTime(*queuedTask.CompletedAt).Format("d M Y")
// 				}).Else("-")
// 				completeddAtTime := lo.IfF(queuedTask.CompletedAt != nil, func() string {
// 					return carbon.CreateFromStdTime(*queuedTask.CompletedAt).ToTimeString()
// 				}).Else("-")

// 				elapsedTime := lo.IfF(queuedTask.StartedAt != nil && queuedTask.CompletedAt != nil, func() string {
// 					return queuedTask.CompletedAt.Sub(*queuedTask.StartedAt).String()
// 				}).Else("-")

// 				return hb.NewTR().Children([]hb.TagInterface{
// 					hb.NewTD().
// 						Child(hb.NewDiv().Text(taskName)).
// 						Child(hb.NewDiv().Text("Alias: ").Text(task.Alias)).
// 						Child(hb.NewDiv().Text("ID: ").Text(queuedTask.ID)),
// 					hb.NewTD().
// 						Child(hb.NewDiv().Text(startedAtDate)).
// 						Child(hb.NewDiv().Text(startedAtTime)).
// 						Style("white-space: nowrap;"),
// 					hb.NewTD().
// 						Child(hb.NewDiv().Text(completeddAtDate)).
// 						Child(hb.NewDiv().Text(completeddAtTime)).
// 						Style("white-space: nowrap;"),
// 					hb.NewTD().
// 						Child(hb.NewDiv().Text(elapsedTime)).
// 						Style("white-space: nowrap;"),
// 					hb.NewTD().
// 						Text(queuedTask.Status),
// 					hb.NewTD().
// 						Style("text-align: center;").
// 						Child(buttonParameters).
// 						Child(buttonDetails).
// 						Child(buttonRequeue).
// 						Child(buttonRestart).
// 						Child(buttonDelete),
// 				})
// 			})),
// 		})

// 	return hb.NewDiv().
// 		ID("QueuedTasksListTable").
// 		Children([]hb.TagInterface{
// 			controller.tableFilter(data),
// 			table,
// 			controller.tablePagination(int(queudTaskCount), pageInt, perPage),
// 		})
// }

// func (controller *queueManagerController) tableFilter(data queueManagerControllerData) hb.TagInterface {
// 	statusList := []map[string]string{
// 		{"id": "", "name": ""},
// 		{"id": models.ACCOUNTING_TRANSACTION_STATUS_DRAFT, "name": "Draft"},
// 		{"id": models.ACCOUNTING_TRANSACTION_STATUS_ACTIVE, "name": "Active"},
// 		{"id": models.ACCOUNTING_TRANSACTION_STATUS_INACTIVE, "name": "Inactive"},
// 		{"id": models.ACCOUNTING_TRANSACTION_STATUS_PENDING, "name": "Pending"},
// 	}

// 	searchButton := hb.NewButton().
// 		Type("submit").
// 		Child(hb.NewI().Class("bi bi-search")).
// 		Class("btn btn-primary w-100 h-100")

// 	period := hb.NewDiv().Class("form-group").Children([]hb.TagInterface{
// 		hb.NewLabel().
// 			HTML("Period").
// 			Style("margin-bottom: 0px;"),
// 		hb.NewDiv().Class("input-group").Children([]hb.TagInterface{
// 			hb.NewInput().
// 				Type(hb.TYPE_DATE).
// 				Name("date_from").
// 				Value(data.dateFrom).
// 				OnChange("FORM_TRANSACTIONS.submit()").
// 				Class("form-control"),
// 			hb.NewSpan().
// 				HTML(" : ").
// 				Class("input-group-text"),
// 			hb.NewInput().
// 				Type(hb.TYPE_DATE).
// 				Name("date_to").
// 				Value(data.dateTo).
// 				OnChange("FORM_TRANSACTIONS.submit()").
// 				Class("form-control"),
// 		}),
// 	})

// 	search := hb.NewDiv().Class("form-group").Children([]hb.TagInterface{
// 		hb.NewLabel().
// 			HTML("Search").
// 			Style("margin-bottom: 0px;"),
// 		hb.NewInput().
// 			Type("search").
// 			Name("search").
// 			Value(data.search).
// 			Class("form-control"),
// 	})

// 	status := hb.NewDiv().Class("form-group").Children([]hb.TagInterface{
// 		hb.NewLabel().
// 			HTML("Status").
// 			Style("margin-bottom: 0px;"),
// 		hb.NewSelect().
// 			Name("status").
// 			Class("form-select").
// 			OnChange("FORM_TRANSACTIONS.submit()").
// 			Children(lo.Map(statusList, func(status map[string]string, index int) hb.TagInterface {
// 				return hb.NewOption().
// 					Value(status["id"]).
// 					HTML(status["name"]).
// 					AttrIf(data.status == status["id"], "selected", "selected")
// 			})),
// 	})

// 	form := hb.NewForm().
// 		ID("FORM_TRANSACTIONS").
// 		Style("display: block").
// 		Method(http.MethodGet).
// 		Children([]hb.TagInterface{
// 			hb.NewDiv().Class("row").Children([]hb.TagInterface{
// 				hb.NewDiv().Class("col-md-3").Children([]hb.TagInterface{
// 					search,
// 				}),
// 				hb.NewDiv().Class("col-md-2").Children([]hb.TagInterface{
// 					status,
// 				}),
// 				hb.NewDiv().Class("col-md-4").Children([]hb.TagInterface{
// 					period,
// 				}),
// 				hb.NewDiv().Class("col-md-1").Children([]hb.TagInterface{
// 					searchButton,
// 				}),
// 			}),
// 		})

// 	return hb.NewDiv().
// 		Class("card bg-light mb-3").
// 		Style("").
// 		Children([]hb.TagInterface{
// 			hb.NewDiv().Class("card-body").Children([]hb.TagInterface{
// 				form,
// 			}),
// 		})
// }

// func (controller *queueManagerController) tablePagination(count int, page int, perPage int) hb.TagInterface {
// 	url := links.NewAdminLinks().Tasks(map[string]string{
// 		"page": "",
// 	})

// 	pagination := bs.Pagination(bs.PaginationOptions{
// 		NumberItems:       count,
// 		CurrentPageNumber: page,
// 		PagesToShow:       20,
// 		PerPage:           perPage,
// 		URL:               url,
// 	})

// 	return hb.NewDiv().
// 		Class(`d-flex justify-content-left mt-5 pagination-primary-soft rounded mb-0`).
// 		HTML(pagination)
// }

// func (controller *queueManagerController) sortableColumnLabel(data queueManagerControllerData, columnName string, columnTableName string) hb.TagInterface {
// 	isSelected := strings.EqualFold(data.sortBy, columnTableName)

// 	direction := lo.If(isSelected && data.sortOrder == "asc", "desc").
// 		Else("asc")

// 	link := links.NewAdminLinks().Tasks(map[string]string{
// 		"page": "0",
// 		"by":   columnTableName,
// 		"sort": direction,
// 	})
// 	return hb.NewHyperlink().
// 		HTML(columnName).
// 		Child(controller.sortingIndicator(columnTableName, data.sortBy, data.sortOrder)).
// 		Href(link)
// }

// func (controller *queueManagerController) prepareData(r *http.Request) (data queueManagerControllerData, errorMessage string) {
// 	data.action = strings.TrimSpace(utils.Req(r, "action", ""))
// 	data.page = strings.TrimSpace(utils.Req(r, "page", "0"))
// 	data.sortOrder = strings.TrimSpace(utils.Req(r, "sort", ""))
// 	data.sortBy = strings.TrimSpace(utils.Req(r, "by", ""))
// 	data.status = strings.TrimSpace(utils.Req(r, "status", ""))
// 	data.search = strings.TrimSpace(utils.Req(r, "search", ""))
// 	data.dateFrom = strings.TrimSpace(utils.Req(r, "date_from", ""))
// 	data.dateTo = strings.TrimSpace(utils.Req(r, "date_to", ""))
// 	data.queueID = strings.TrimSpace(utils.Req(r, "queue_id", ""))
// 	data.taskID = strings.TrimSpace(utils.Req(r, "task_id", ""))

// 	if !lo.Contains([]string{models.ASC, models.DESC}, data.sortOrder) {
// 		data.sortOrder = models.DESC
// 	}

// 	if !lo.Contains([]string{"started_at", "ended_at", "id", "task_id", "elapsed", "status"}, data.sortBy) {
// 		data.sortBy = models.COLUMN_CREATED_AT
// 	}

// 	cfmt.Errorln("queueManagerController > prepareData > ", data.sortBy, data.sortOrder)

// 	return data, ""
// }

// func (controller *queueManagerController) sortingIndicator(columnName string, sortByColumnName string, sortOrder string) hb.TagInterface {
// 	isSelected := strings.EqualFold(sortByColumnName, columnName)

// 	direction := lo.If(isSelected && sortOrder == "asc", "up").
// 		ElseIf(isSelected && sortOrder == "desc", "down").
// 		Else("none")

// 	sortingIndicator := hb.NewSpan().
// 		Class("sorting").
// 		HTMLIf(direction == "up", "&#8595;").
// 		HTMLIf(direction == "down", "&#8593;").
// 		HTMLIf(direction != "down" && direction != "up", "")

// 	return sortingIndicator
// }
