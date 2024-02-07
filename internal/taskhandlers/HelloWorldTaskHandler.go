package taskhandlers

import (
	"project/config"

	"github.com/gouniverse/taskstore"
)

func NewHelloWorldTaskHandler() *helloWorldTaskHandler {
	return &helloWorldTaskHandler{}
}

type helloWorldTaskHandler struct {
	taskstore.TaskHandlerBase
}

var _ taskstore.TaskHandlerInterface = (*helloWorldTaskHandler)(nil) // verify it extends the task interface

func (handler *helloWorldTaskHandler) Alias() string {
	return "helloWorldTaskHandler"
}

func (handler *helloWorldTaskHandler) Title() string {
	return "Hello World"
}

func (handler *helloWorldTaskHandler) Description() string {
	return "Say hello world"
}

func (handler *helloWorldTaskHandler) Enqueue() (task *taskstore.Queue, err error) {
	return config.TaskStore.TaskEnqueueByAlias(handler.Alias(), map[string]any{})
}

func (handler *helloWorldTaskHandler) Handle() bool {
	if !handler.HasQueuedTask() && handler.GetParam("enqueue") == "yes" {
		_, err := handler.Enqueue()

		if err != nil {
			handler.LogError("Error enqueuing task: " + err.Error())
		} else {
			handler.LogSuccess("Task enqueued.")
		}

		return true
	}

	handler.LogInfo("Hello World!")
	return true
}
