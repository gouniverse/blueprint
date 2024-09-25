package admin

import (
	"net/http"
	"project/config"
	"strings"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/maputils"
	"github.com/gouniverse/utils"
)

func (controller *queueManagerController) onModalQueuedTaskRequeueSubmitted(r *http.Request) string {
	taskID := strings.TrimSpace(utils.Req(r, "task_id", ""))
	taskParameters := strings.TrimSpace(utils.Req(r, "task_parameters", ""))

	if taskID == "" {
		return hb.NewSwal(hb.SwalOptions{Title: "Error", Text: "Task is required"}).ToHTML()
	}

	if taskParameters == "" {
		taskParameters = "{}"
	}

	if !utils.IsJSON(taskParameters) {
		return hb.NewSwal(hb.SwalOptions{Icon: "error", Title: "Error", Text: "Task Parameters is not valid JSON"}).ToHTML()
	}

	task := config.TaskStore.TaskFindByID(taskID)
	if task == nil {
		return hb.NewSwal(hb.SwalOptions{Title: "Error", Text: "Task not found"}).ToHTML()
	}

	taskParametersAny, err := utils.FromJSON(taskParameters, map[string]interface{}{})

	if err != nil {
		config.LogStore.ErrorWithContext("At adminTasks > onModalTaskEnqueueSubmitted", err.Error())
		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be enqueued").ToHTML()
	}

	taskParametersMap := maputils.AnyToMapStringAny(taskParametersAny)

	_, err = config.TaskStore.TaskEnqueueByAlias(task.Alias, taskParametersMap)
	if err != nil {
		config.LogStore.ErrorWithContext("At adminTasks > onModalTaskEnqueueSubmitted", err.Error())
		return hb.NewDiv().Class("alert alert-danger").Text("Task failed to be enqueued").ToHTML()
	}

	response := hb.NewSwal(hb.SwalOptions{Icon: "success", Title: "Success", Text: "Task enqueued successfully"}).ToHTML()

	response += hb.NewScript(`setTimeout(() => {window.location.href = window.location.href;}, 3000);`).ToHTML()

	return response

}
