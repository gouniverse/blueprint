package cli

import (
	"errors"
	"project/app/routes"
	"project/config"
	"project/internal/cmds"

	"github.com/gouniverse/router"
	"github.com/mingrammer/cfmt"
)

// executeCliCommand executes a CLI command
//
// The command can be one of the following:
// - task <alias> <arguments>
// - job <arguments>
// - routes list
//
// Business logic:
//
// 1. First, it logs the command being executed, so it's obvious what's going on.
// 2. It checks if there are at least two arguments, and appends "list" to the arguments if not.
// 3. It gets the first and second arguments.
// 4. If the first argument is "task", it executes the task with the second argument any additional arguments.
// 5. If the first argument is "job", it executes the job with any additional arguments.
// 6. If the first argument is "routes" and the second argument is "list" it lists all the routes.
// 7. Otherwise, it prints a warning that the command is unrecognized.
//
// Parameters:
// - args []string : The command line arguments.
//
// Returns:
// - none
func ExecuteCliCommand(args []string) error {
	cfmt.Infoln("Executing command: ", args)
	if len(args) < 2 {
		args = append(args, "list")
	}

	firstArg := args[0]
	secondArg := args[1]

	// Is it a task?
	if firstArg == "task" {
		if config.TaskStore == nil {
			cfmt.Errorln("TaskStore is nil")
			return errors.New("task store is nil")
		}

		config.TaskStore.TaskExecuteCli(secondArg, args[2:])

		return nil
	}

	// Is it a job?
	if firstArg == "job" {
		cmds.ExecuteJob(args[2:])
		return nil
	}

	// Is it a route list?
	if firstArg == "routes" && secondArg == "list" {
		m, r := routes.RoutesList()
		router.List(m, r)
		return nil
	}

	cfmt.Warning("Unrecognized command: ", firstArg)
	return nil
}
