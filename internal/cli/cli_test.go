package cli

import (
	// Import errors package
	"fmt" // Import fmt package
	"os"
	"testing"

	"project/config"
	"project/internal/testutils"
)

func TestExecuteCliCommand_NilTaskStore(t *testing.T) {
	testutils.Setup()

	// Test task execution with TaskStore nil
	os.Args = []string{"main", "task", "testTask"}
	config.TaskStore = nil // Ensure TaskStore is nil for this specific test case
	err := ExecuteCliCommand(os.Args[1:])

	if err == nil {
		t.Errorf("Expected error when TaskStore is nil, but got nil")
	} else {
		expectedErrMsg := "task store is nil"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got: %v", expectedErrMsg, err)
		}
	}
}

func TestExecuteCliCommand_TaskExecution(t *testing.T) {
	testutils.Setup() // Sets up necessary config, including potentially a mock TaskStore

	// Test task execution with TaskStore not nil
	os.Args = []string{"main", "task", "testTask"}
	// Ensure TaskStore is not nil (testutils.Setup should handle this, or mock it)
	if config.TaskStore == nil {
		t.Skip("Skipping test: TaskStore is nil, cannot test task execution.") // Or setup a mock TaskStore here
	}
	err := ExecuteCliCommand(os.Args[1:])
	if err != nil {
		t.Errorf("Expected no error during task execution, got: %v", err)
	}
	// Add assertions here to check if the mock task was actually called if applicable
}

func TestExecuteCliCommand_JobExecution(t *testing.T) {
	testutils.Setup()

	// Test job execution
	os.Args = []string{"main", "job", "testJob"}
	err := ExecuteCliCommand(os.Args[1:])
	if err != nil {
		t.Errorf("Expected no error during job execution, got: %v", err)
	}
	// Add assertions here to check if the mock job was actually called if applicable
}

func TestExecuteCliCommand_RoutesList(t *testing.T) {
	testutils.Setup()

	// Test routes list
	os.Args = []string{"main", "routes", "list"}
	err := ExecuteCliCommand(os.Args[1:])
	if err != nil {
		t.Errorf("Expected no error during routes list execution, got: %v", err)
	}
	// Add assertions here to check if the router.List was called, maybe by capturing stdout
}

func TestExecuteCliCommand_UnrecognizedCommand(t *testing.T) {
	testutils.Setup()

	// Test unrecognized command
	commandName := "unknownCommand"
	os.Args = []string{"main", commandName}
	err := ExecuteCliCommand(os.Args[1:])

	// Expect an error for unrecognized commands
	if err == nil {
		t.Errorf("Expected an error for unrecognized command '%s', but got nil", commandName)
	} else {
		// Optional: Check if the error message is what you expect
		expectedErrMsg := fmt.Sprintf("unrecognized command: %s", commandName)
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%v'", expectedErrMsg, err)
		}
	}
}

func TestExecuteCliCommand_NoCommand(t *testing.T) {
	testutils.Setup()

	// Test with no command provided
	os.Args = []string{"main"} // Only program name
	err := ExecuteCliCommand(os.Args[1:])

	if err == nil {
		t.Errorf("Expected an error when no command is provided, but got nil")
	} else {
		expectedErrMsg := "no command provided"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%v'", expectedErrMsg, err)
		}
	}
}

func TestExecuteCliCommand_TaskMissingAlias(t *testing.T) {
	testutils.Setup()

	// Test task command without alias
	os.Args = []string{"main", CommandTask}
	err := ExecuteCliCommand(os.Args[1:])

	if err == nil {
		t.Errorf("Expected an error when task alias is missing, but got nil")
	} else {
		expectedErrMsg := fmt.Sprintf("missing task alias for command '%s'", CommandTask)
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%v'", expectedErrMsg, err)
		}
	}
}

func TestExecuteCliCommand_RoutesMissingSubcommand(t *testing.T) {
	testutils.Setup()

	// Test routes command without subcommand
	os.Args = []string{"main", CommandRoutes}
	err := ExecuteCliCommand(os.Args[1:])

	if err == nil {
		t.Errorf("Expected an error when routes subcommand is missing, but got nil")
	} else {
		expectedErrMsg := fmt.Sprintf("invalid or missing subcommand for '%s'. Use '%s %s'", CommandRoutes, CommandRoutes, SubcommandList)
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%v'", expectedErrMsg, err)
		}
	}
}

func TestExecuteCliCommand_RoutesInvalidSubcommand(t *testing.T) {
	testutils.Setup()

	// Test routes command with invalid subcommand
	os.Args = []string{"main", CommandRoutes, "invalidsub"}
	err := ExecuteCliCommand(os.Args[1:])

	if err == nil {
		t.Errorf("Expected an error when routes subcommand is invalid, but got nil")
	} else {
		expectedErrMsg := fmt.Sprintf("invalid or missing subcommand for '%s'. Use '%s %s'", CommandRoutes, CommandRoutes, SubcommandList)
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%v'", expectedErrMsg, err)
		}
	}
}
