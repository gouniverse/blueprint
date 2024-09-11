package main

import (
	"log"
	"os"

	"project/config"
	"project/internal/cmds"
	"project/internal/routes"
	"project/internal/scheduler"
	"project/internal/tasks"
	"project/internal/widgets"
	"project/models"

	"github.com/gouniverse/router"
	"github.com/gouniverse/webserver"
	"github.com/mingrammer/cfmt"
)

// main starts the application
//
// Business Logic:
// 1. Initialize the environment
// 2. Defer Closing the database
// 3. Initialize the models
// 4. Register the task handlers
// 5. Executes the command if provided
// 6. Initialize the task queue
// 7. Initialize the scheduler
// 8. Starts the cache expiration goroutine
// 9. Starts the session expiration goroutine
// 10. Adds CMS shortcodes
// 11. Starts the web server
//
// Parameters:
// - none
//
// Returns:
// - none
func main() {
	config.Initialize()    // 1. Initialize the environment
	defer closeResources() // 2. Defer Closing the database
	models.Initialize()    // 3. Initialize the models
	tasks.RegisterTasks()  // 4. Register the task handlers

	if isCliMode() {
		if len(os.Args) < 2 {
			return
		}
		executeCliCommand(os.Args[1:]) // 5. Execute the command
		return
	}

	startBackgroundProcesses()
	startWebServer() // 11. Start the web server
}

func closeResources() {
	if err := config.Database.DB().Close(); err != nil {
		cfmt.Errorf("Failed to close database connection: %v", err)
	}
}

func isCliMode() bool {
	return len(os.Args) > 1
}

func startBackgroundProcesses() {
	go config.TaskStore.QueueRunGoroutine(10, 2)    // 6. Initialize the task queue
	scheduler.StartAsync()                          // 7. Initialize the scheduler
	go config.CacheStore.ExpireCacheGoroutine()     // 8. Initialize the cache expiration goroutine
	go config.SessionStore.ExpireSessionGoroutine() // 9. Initialize the session expiration goroutine
	widgets.CmsAddShortcodes()                      // 10. Add CMS shortcodes
}

// executeCommand executes a CLI command
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
func executeCliCommand(args []string) {
	cfmt.Infoln("Executing command: ", args)
	if len(args) < 2 {
		args = append(args, "list")
	}

	firstArg := args[0]
	secondArg := args[1]

	// Is it a task?
	if firstArg == "task" {
		config.TaskStore.TaskExecuteCli(secondArg, args[2:])
		return
	}

	// Is it a job?
	if firstArg == "job" {
		cmds.ExecuteJob(args[2:])
		return
	}

	// Is it a route list?
	if firstArg == "routes" && secondArg == "list" {
		m, r := routes.RoutesList()
		router.List(m, r)
		return
	}

	cfmt.Warning("Unrecognized command: ", firstArg)
}

// startServer starts the web server at the specified host and port and listens
// for incoming requests.
//
// Parameters:
// - none
//
// Returns:
// - none
func startWebServer() {
	addr := config.WebServerHost + ":" + config.WebServerPort
	cfmt.Infoln("Starting server on " + config.WebServerHost + ":" + config.WebServerPort + " ...")
	cfmt.Infoln("APP URL: " + config.AppUrl + " ...")

	config.WebServer = webserver.New(addr, routes.Routes().ServeHTTP)

	if err := config.WebServer.ListenAndServe(); err != nil {
		if config.AppEnvironment == config.APP_ENVIRONMENT_TESTING {
			cfmt.Errorln(err)
		} else {
			log.Fatal(err)
		}
	}
}
