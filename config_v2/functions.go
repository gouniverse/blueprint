package config_v2

import "context"

// ToContext adds the configuration to the context
func ToContext(ctx context.Context, config *Config) context.Context {
	return context.WithValue(ctx, configContextKey{}, config)
}

// FromContext retrieves the configuration from the context
func FromContext(ctx context.Context) *Config {
	if cfg, ok := ctx.Value(configContextKey{}).(*Config); ok {
		return cfg
	}
	return nil
}
