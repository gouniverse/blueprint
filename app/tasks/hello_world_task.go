package tasks

import (
	"errors"
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
	return "HelloWorldTask"
}

func (handler *helloWorldTask) Title() string {
	return "Hello World"
}

func (handler *helloWorldTask) Description() string {
	return "Say hello world"
}

func (handler *helloWorldTask) Enqueue() (task taskstore.QueueInterface, err error) {
	if config.TaskStore == nil {
		return nil, errors.New("task store is nil")
	}
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
