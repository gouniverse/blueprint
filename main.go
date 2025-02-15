package main

import (
	"os"
	"project/config"
	"project/internal/cli"
	"project/internal/middlewares"
	"project/internal/scheduler"
	"project/internal/tasks"
	"project/internal/widgets"

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
	tasks.RegisterTasks()  // 3. Register the task handlers

	if isCliMode() {
		if len(os.Args) < 2 {
			return
		}
		cli.ExecuteCliCommand(os.Args[1:]) // 4. Execute the command
		return
	}

	startBackgroundProcesses()
	StartWebServer() // 5. Start the web server
}

// closeResources closes the database connection if it exists.
//
// Parameters:
// - none
//
// Returns:
// - none
func closeResources() {
	if config.Database == nil {
		return
	}

	if err := config.Database.DB().Close(); err != nil {
		cfmt.Errorf("Failed to close database connection: %v", err)
	}
}

// isCliMode checks if the application is running in CLI mode.
//
// Parameters:
// - none
//
// Returns:
// - bool: true if the application is running in CLI mode, false otherwise.
func isCliMode() bool {
	return len(os.Args) > 1
}

// startBackgroundProcesses starts the background processes.
//
// Parameters:
// - none
//
// Returns:
// - none
func startBackgroundProcesses() {
	if config.TaskStore != nil {
		go config.TaskStore.QueueRunGoroutine(10, 2) // Initialize the task queue
	}

	scheduler.StartAsync() // Initialize the scheduler

	go config.CacheStore.ExpireCacheGoroutine() // Initialize the cache expiration goroutine

	if config.SessionStore != nil {
		go config.SessionStore.SessionExpiryGoroutine() // Initialize the session expiration goroutine
	}

	middlewares.CmsAddMiddlewares() // Add CMS middlewares
	widgets.CmsAddShortcodes()      // Add CMS shortcodes
}
