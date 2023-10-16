package main

import (
	"log"
	"os"

	"project/config"
	"project/internal/cmds"
	"project/internal/routes"
	"project/internal/server"
	"project/internal/taskhandlers"
	"project/models"

	"github.com/gouniverse/router"
	"github.com/mingrammer/cfmt"
)

func main() {
	cfmt.Infoln("Initializing configuration ...")
	config.Initialize()                // 1. Initialize the environment
	defer config.Database.DB().Close() // 2. Defer Closing the database

	models.Initialize()    // 3. Initialize the models
	registerTaskHandlers() // 4. Register the task handlers
	// jobs.Initialize()   // 5. Initialize the jobs

	// If there are arguments, run the command interface
	if len(os.Args) > 1 {
		executeCommand(os.Args[1:]) // 6. Execute the command
		return
	}

	// jobs.RunScheduler() // 7. Run the scheduler

	startServer() // 8. Start the server
}

func registerTaskHandlers() {
	cfmt.Infoln("Registering task handlers ...")
	config.Cms.TaskStore.TaskHandlerAdd(taskhandlers.NewHelloWorldTaskHandler(), true)
}

// executeCommand executes a command
func executeCommand(args []string) {
	cfmt.Infoln("Executing command: ", args)
	if len(args) < 2 {
		args = append(args, "list")
	}

	firstArg := args[0]
	secondArg := args[1]
	if firstArg == "task" {
		cmds.ExecuteCommand(secondArg, args[2:])
		return
	}

	if firstArg == "job" {
		cmds.ExecuteJob(args[2:])
		return
	}

	if firstArg == "routes" && secondArg == "list" {
		cfmt.Warning("Unrecognized command: ", firstArg)
		m, r := routes.RoutesList()
		router.List(m, r)
		return
	}
}

func startServer() {
	addr := config.ServerHost + ":" + config.ServerPort
	cfmt.Infoln("Starting server on " + config.ServerHost + ":" + config.ServerPort + " ...")
	cfmt.Infoln("APP URL: " + config.AppUrl + " ...")
	config.WebServer = server.NewServer(addr, routes.Routes().ServeHTTP)
	if err := config.WebServer.ListenAndServe(); err != nil {
		if config.AppEnvironment == config.APP_ENVIRONMENT_TESTING {
			cfmt.Errorln(err)
		} else {
			log.Fatal(err)
		}
	}
}
