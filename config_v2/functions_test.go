package config_v2

import "testing"

func TestIsDebugEnabled(t *testing.T) {
	cfg, err := TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.GetDebug() && !cfg.IsDebugEnabled() {
		t.Fatal(`Must be in debug, if debug enabled (1)`)
	}

	cfg.SetDebug(false)

	if cfg.Debug && !cfg.IsDebugEnabled() {
		t.Fatal(`Must not be in debug, if debug enabled (2)`)
	}
}

func TestIsEnvDevelopment(t *testing.T) {
	cfg, err := TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.GetAppEnvironment() != APP_ENVIRONMENT_DEVELOPMENT && cfg.IsEnvDevelopment() {
		t.Fatal(`Must not be in development`)
	}

	cfg.SetAppEnvironment(APP_ENVIRONMENT_DEVELOPMENT)

	if !cfg.IsEnvDevelopment() {
		t.Fatal(`Must be in development`)
	}

	if cfg.GetAppEnvironment() != APP_ENVIRONMENT_DEVELOPMENT && cfg.IsEnvDevelopment() {
		t.Fatal(`Must not be in development`)
	}
}

func TestIsEnvProduction(t *testing.T) {
	cfg, err := TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.IsEnvProduction() {
		t.Fatal(`Must not be in production`)
	}

	if cfg.GetAppEnvironment() != APP_ENVIRONMENT_PRODUCTION && cfg.IsEnvProduction() {
		t.Fatal(`Must not be in production`)
	}

	cfg.SetAppEnvironment(APP_ENVIRONMENT_PRODUCTION)

	if !cfg.IsEnvProduction() {
		t.Fatal(`Must be in production`)
	}

	if cfg.GetAppEnvironment() != APP_ENVIRONMENT_PRODUCTION && cfg.IsEnvProduction() {
		t.Fatal(`Must not be in production`)
	}
}

func TestIsEnvLocal(t *testing.T) {
	cfg, err := TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.IsEnvLocal() {
		t.Fatal(`Must not be in local`)
	}

	cfg.SetAppEnvironment(APP_ENVIRONMENT_LOCAL)

	if !cfg.IsEnvLocal() {
		t.Fatal(`Must be in local`)
	}

	cfg, err = TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.IsEnvLocal() {
		t.Fatal(`Must not be in local`)
	}
}

func TestIsEnvTesting(t *testing.T) {
	cfg, err := TestsConfigureAndInitialize()

	cfg.AppEnvironment = APP_ENVIRONMENT_LOCAL

	if err != nil {
		t.Fatal(err)
	}

	if cfg.IsEnvTesting() {
		t.Fatal(`Must not be in testing`)
	}

	cfg.SetAppEnvironment(APP_ENVIRONMENT_TESTING)

	if !cfg.IsEnvTesting() {
		t.Fatal(`Must be in testing`)
	}

	cfg, err = TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if !cfg.IsEnvTesting() {
		t.Fatal(`Must not be in testing`)
	}
}
