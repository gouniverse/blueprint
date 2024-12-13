package cmds

import (

	// "project/internal/jobs"

	// "project/jobs"

	"project/config"

	"github.com/gouniverse/taskstore"
	"github.com/gouniverse/utils"
	"github.com/mingrammer/cfmt"
	"github.com/samber/lo"
)

// ExecuteJob executes a job (queued task) with the given arguments.
//
// Example:
// go run . job run --task_id=20231008040147830106 --force=yes
//
// Args: an array of strings representing the arguments for the job.
// Return type: None.
func ExecuteJob(args []string) {
	name := "No name"
	argumentsMap := utils.ArgsToMap(args)
	cfmt.Infoln("Executing job: ", name, " with arguments: ", argumentsMap)

	queuedTaskID := lo.ValueOr(argumentsMap, "task_id", "")
	force := lo.ValueOr(argumentsMap, "force", "")

	if queuedTaskID == "" {
		cfmt.Errorln("Task ID is required")
		return
	}

	if config.TaskStore == nil {
		cfmt.Errorln("TaskStore is nil")
		return
	}

	queuedTask, err := config.TaskStore.QueueFindByID(queuedTaskID)

	if err != nil {
		cfmt.Errorln("Task not found: ", queuedTaskID)
		return
	}

	if queuedTask == nil {
		cfmt.Errorln("Task not found: ", queuedTaskID)
		return
	}

	if queuedTask.Status() == taskstore.QueueStatusRunning {
		cfmt.Errorln("Task is currently running: ", queuedTaskID, "Aborted")
		return
	}

	if force != "yes" && queuedTask.Status() != taskstore.QueueStatusQueued {
		cfmt.Errorln("Task is not queued: ", queuedTaskID, " . You can use the --force=yes option to force the execution of the job. Aborted")
		return
	}

	isOK, err := config.TaskStore.QueuedTaskProcess(queuedTask)

	if err != nil {
		cfmt.Errorln("Error processing task: ", queuedTaskID, " ", err.Error())
		return
	}

	if isOK {
		cfmt.Infoln("Job: ", queuedTaskID, " run OK")
	} else {
		cfmt.Errorln("Job: ", queuedTaskID, " run failed")
	}
}
