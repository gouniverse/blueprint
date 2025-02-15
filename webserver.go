package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"project/config"
	"project/internal/routes"

	"github.com/gouniverse/webserver"
	"github.com/mingrammer/cfmt"
)

var shutdownChan = make(chan os.Signal, 1)

// StartWebServerbserver starts the web server at the specified host and port and listens
// for incoming requests.
//
// Parameters:
// - none
//
// Returns:
// - none
func StartWebServer() {
	addr := config.WebServerHost + ":" + config.WebServerPort
	cfmt.Infoln("Starting server on " + config.WebServerHost + ":" + config.WebServerPort + " ...")
	cfmt.Infoln("APP URL: " + config.AppUrl + " ...")

	config.WebServer = webserver.New(addr, routes.Routes().ServeHTTP)

	// Register the shutdown signal
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		if err := config.WebServer.ListenAndServe(); err != nil {
			if config.AppEnvironment == config.APP_ENVIRONMENT_TESTING {
				cfmt.Errorln(err)
			} else {
				log.Fatal(err)
			}
		}
	}()

	// Wait for a shutdown signal
	cfmt.Infoln("Server is running, press Ctrl+C to stop it.")
	select {
	case sig := <-shutdownChan:
		cfmt.Infoln("Received signal:", sig)
	}
	cfmt.Infoln("Shutting down server...")
	config.WebServer.Shutdown(context.Background())
	if config.AppEnvironment != config.APP_ENVIRONMENT_TESTING {
		os.Exit(0)
	}
}
