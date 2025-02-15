package main

import (
	"fmt"
	"net/http"
	"sync"
	"syscall"
	"testing"
	"time"

	"project/config"
	"project/internal/testutils"

	"github.com/stretchr/testify/assert"
)

func TestStartWebServer(t *testing.T) {
	testutils.Setup()

	var wg sync.WaitGroup
	wg.Add(1)

	// Start the web server in a goroutine
	go func() {
		defer wg.Done()
		StartWebServer()
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Check if the server is running
	url := fmt.Sprintf("http://%s:%s", config.WebServerHost, config.WebServerPort)
	resp, err := http.Get(url)
	assert.NoError(t, err, "Failed to make a request to the server")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Server should return status OK")

	// Send a shutdown signal to the shutdownChan
	shutdownChan <- syscall.SIGTERM

	// Wait for the server to shut down
	wg.Wait()

	// Check if the server is shut down
	_, err = http.Get(url)
	assert.Error(t, err, "Server should be shut down")
}
