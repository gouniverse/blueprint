package app

import (
	"context"

	"project/config_v2"
)

// GetConfig retrieves the configuration from the context.
// This helps transition from the global configuration system
// to a context-based approach for better testability and parallel test execution.
func GetConfig(ctx context.Context) interface{} {
	return config_v2.FromContext(ctx)
}
