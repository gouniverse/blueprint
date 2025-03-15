package tasks

import (
	"os"

	"github.com/gouniverse/envenc"
	"github.com/gouniverse/taskstore"
	"github.com/mingrammer/cfmt"
)

// ===================================================================
// envencTask
// ===================================================================
// Adds commnds for the .env.vault file
// ==================================================================
// Example:
// - go run main.go task EnvEncTask init .env.vault
// - go run main.go task EnvEncTask key-set .env.vault
// - go run main.go task EnvEncTask key-list .env.vault
// - go run main.go task EnvEncTask key-remove .env.vault
// ==================================================================
type envencTask struct {
	taskstore.TaskHandlerBase
}

var _ taskstore.TaskHandlerInterface = (*envencTask)(nil) // verify it extends the task interface

func NewEnvencTask() *envencTask {
	return &envencTask{}
}

func (handler *envencTask) Alias() string {
	return "EnvEncTask"
}

func (handler *envencTask) Title() string {
	return "EnvEnc"
}

func (handler *envencTask) Description() string {
	return "Utilities for the .env.vault file"
}

func (handler *envencTask) Handle() bool {
	if len(os.Args) < 3 {
		cfmt.Errorln("Usage: go run main.go task EnvEncTask <command> .env.vault")
		return false
	}

	envenc.NewCli().Run(os.Args[2:])
	return true
}
