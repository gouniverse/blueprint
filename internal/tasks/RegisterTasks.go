package tasks

import (
	"project/config"

	"github.com/gouniverse/taskstore"
)

// RegisterTasks registers the task handlers to the task store
//
// Parameters:
// - none
//
// Returns:
// - none
func RegisterTasks() {
	tasks := []taskstore.TaskHandlerInterface{
		NewBlindIndexRebuildTask(),
		NewEnvencTask(),
		NewHelloWorldTask(),
		NewStatsVisitorEnhanceTask(),
	}

	for _, task := range tasks {
		err := config.TaskStore.TaskHandlerAdd(task, true)

		if err != nil {
			config.LogStore.ErrorWithContext("At registerTaskHandlers", "Error registering task: "+task.Alias()+" - "+err.Error())
		}
	}
}
