package testutils

import (
	"project/config"
	"project/config_v2"
	//smtpmock "github.com/mocktools/go-smtp-mock"
)

func Setup() {
	if config.Database != nil {
		config.Database.Exec(`PRAGMA writable_schema = 1;
		DELETE FROM sqlite_master;
		PRAGMA writable_schema = 0;
		VACUUM;
		PRAGMA integrity_check;`)
	}

	config.TestsConfigureAndInitialize()
	// var errAuthInit error
	// config.Auth, errAuthInit = authentication.SetupAuth()
	// if errAuthInit != nil {
	// 	config.LogStore.Error("Auth Initialization Failed: " + errAuthInit.Error())
	// 	log.Panicln("Auth Initialization Failed: " + errAuthInit.Error())
	// }
	// database.Initialize()
	// authentication.Initialize()
	// jobs.Initialize()
	// setupMailServer()
	// tasks.RegisterTasks()
}

func SetupV2SetEnvironmentVariablesOnly() {
	config_v2.TestsSetEnvironmentVariables()
}

func SetupV2() (*config_v2.Config, error) {
	cfg, err := config_v2.TestsConfigureAndInitialize()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// func setupMailServer() {
// 	mailServer := smtpmock.New(smtpmock.ConfigurationAttr{
// 		LogToStdout:       false, // enable if you have errors sending emails
// 		LogServerActivity: true,
// 		PortNumber:        32435,
// 		HostAddress:       "127.0.0.1",
// 	})

// 	if err := mailServer.Start(); err != nil {
// 		fmt.Println(err)
// 	}
// }
