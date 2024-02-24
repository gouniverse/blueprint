package tasks

import (
	"project/config"

	"github.com/gouniverse/taskstore"
)

func NewHelloWorldTask() *helloWorldTask {
	return &helloWorldTask{}
}

type helloWorldTask struct {
	taskstore.TaskHandlerBase
}

var _ taskstore.TaskHandlerInterface = (*helloWorldTask)(nil) // verify it extends the task interface

func (handler *helloWorldTask) Alias() string {
	return "helloWorldTask"
}

func (handler *helloWorldTask) Title() string {
	return "Hello World"
}

func (handler *helloWorldTask) Description() string {
	return "Say hello world"
}

func (handler *helloWorldTask) Enqueue() (task *taskstore.Queue, err error) {
	return config.TaskStore.TaskEnqueueByAlias(handler.Alias(), map[string]any{})
}

func (handler *helloWorldTask) Handle() bool {
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
