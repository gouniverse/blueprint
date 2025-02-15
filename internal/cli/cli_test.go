package cli

import (
	"os"
	"testing"

	"project/config"
	"project/internal/testutils"
)

func TestExecuteCliCommand_NilTaskStore(t *testing.T) {
	testutils.Setup()

	// Test task execution with TaskStore nil
	os.Args = []string{"main", "task", "testTask"}
	config.TaskStore = nil
	err := ExecuteCliCommand(os.Args[1:])

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "task store is nil" {
		t.Errorf("Expected error 'task store is nil', got: %v", err)
	}
}

func TestExecuteCliCommand_TaskExecution(t *testing.T) {
	testutils.Setup()

	// Test task execution with TaskStore not nil
	os.Args = []string{"main", "task", "testTask"}
	err := ExecuteCliCommand(os.Args[1:])
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestExecuteCliCommand_JobExecution(t *testing.T) {
	testutils.Setup()

	// Test job execution
	os.Args = []string{"main", "job", "testJob"}
	err := ExecuteCliCommand(os.Args[1:])
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestExecuteCliCommand_RoutesList(t *testing.T) {
	testutils.Setup()

	// Test routes list
	os.Args = []string{"main", "routes", "list"}
	err := ExecuteCliCommand(os.Args[1:])
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestExecuteCliCommand_UnrecognizedCommand(t *testing.T) {
	testutils.Setup()

	// Test unrecognized command
	os.Args = []string{"main", "unknownCommand"}
	err := ExecuteCliCommand(os.Args[1:])
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
