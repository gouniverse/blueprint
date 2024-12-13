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
		NewCmsTransferTask(),
		NewEmailToAdminTaskHandler(),
		NewEmailToAdminOnNewContactFormSubmittedTaskHandler(),
		NewEmailToAdminOnNewUserRegisteredTaskHandler(),
		NewEnvencTask(),
		NewHelloWorldTask(),
		NewStatsVisitorEnhanceTask(),
	}

	if config.TaskStore == nil {
		return
	}

	for _, task := range tasks {
		err := config.TaskStore.TaskHandlerAdd(task, true)

		if err != nil {
			config.Logger.Error("At registerTaskHandlers", "error", "Error registering task: "+task.Alias()+" - "+err.Error())
		}
	}
}
