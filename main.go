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
	"github.com/gouniverse/taskstore"
	"github.com/gouniverse/webserver"
	"github.com/mingrammer/cfmt"
)

func main() {
	cfmt.Infoln("Initializing configuration ...")
	config.Initialize()                // 1. Initialize the environment
	defer config.Database.DB().Close() // 2. Defer Closing the database

	models.Initialize()    // 3. Initialize the models
	registerTaskHandlers() // 4. Register the task handlers

	// If there are arguments, run the command interface
	if len(os.Args) > 1 {
		executeCliCommand(os.Args[1:]) // 5. Execute the command
		return
	}

	queueInitialize()      // 6. Initialize the task queue
	scheduler.StartAsync() // 7. Initialize the scheduler

	go config.CacheStore.ExpireCacheGoroutine()
	go config.SessionStore.ExpireSessionGoroutine()

	widgets.CmsAddShortcodes()

	startServer() // 8. Start the server
}

func queueInitialize() {
	go config.TaskStore.QueueRunGoroutine(10, 2)
}

func registerTaskHandlers() {
	cfmt.Infoln("Registering task handlers ...")
	tasks := []taskstore.TaskHandlerInterface{
		tasks.NewEnvencTask(),
		tasks.NewHelloWorldTask(),
	}

	for _, task := range tasks {
		err := config.TaskStore.TaskHandlerAdd(task, true)

		if err != nil {
			config.LogStore.ErrorWithContext("At registerTaskHandlers", "Error registering task: "+task.Alias()+" - "+err.Error())
		}
	}
}

// executeCommand executes a command
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
		cfmt.Warning("Unrecognized command: ", firstArg)
		m, r := routes.RoutesList()
		router.List(m, r)
		return
	}
}

func startServer() {
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
