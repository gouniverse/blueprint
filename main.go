package main

import (
	"log"
	"net/http"
	"os"

	"project/config"
	"project/internal/cmds"
	"project/internal/routes"
	"project/internal/scheduler"
	"project/internal/server"
	"project/internal/tasks"
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

	// If there are arguments, run the command interface
	if len(os.Args) > 1 {
		executeCliCommand(os.Args[1:]) // 5. Execute the command
		return
	}

	queueInitialize()      // 6. Initialize the task queue
	scheduler.StartAsync() // 7. Initialize the scheduler

	go config.CacheStore.ExpireCacheGoroutine()
	go config.SessionStore.ExpireSessionGoroutine()

	startServer() // 8. Start the server
}

func queueInitialize() {
	go config.TaskStore.QueueRunGoroutine(10, 2)
}

func registerTaskHandlers() {
	cfmt.Infoln("Registering task handlers ...")
	config.TaskStore.TaskHandlerAdd(tasks.NewHelloWorldTask(), true)
	config.TaskStore.TaskHandlerAdd(tasks.NewEnvencTask(), true)
}

// executeCommand executes a command
func executeCliCommand(args []string) {
	cfmt.Infoln("Executing command: ", args)
	if len(args) < 2 {
		args = append(args, "list")
	}

	firstArg := args[0]
	secondArg := args[1]
	if firstArg == "task" {
		config.TaskStore.TaskExecuteCli(secondArg, args[2:])
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

func cmsAddShortcodes() {
	shortcodes := map[string]func(*http.Request, string, map[string]string) string{
		//"x-authenticated":   widgets.NewAuthenticatedWidget().Render,
		//"x-unauthenticated": widgets.NewUnauthenticatedWidget().Render,
		//"x-print":                    widgets.NewPrintWidget().Render,
	}

	config.Cms.ShortcodesAdd(shortcodes)
}
