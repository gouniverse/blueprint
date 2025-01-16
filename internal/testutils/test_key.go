package testutils

import (
	"project/config"

	"github.com/gouniverse/utils"
)

// TestKey is a pseudo secret test key used for testing specific unit cases
//
//	where a secret key is required but not available in the testing environment
func TestKey() string {
	return utils.StrToMD5Hash(config.DbDriver + config.DbHost + config.DbPort + config.DbName + config.DbUser + config.DbPass)
}
