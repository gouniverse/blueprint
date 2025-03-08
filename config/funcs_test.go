package config

import "testing"

func TestIsDebugEnabled(t *testing.T) {
	if Debug && !IsDebugEnabled() {
		t.Fatal(`Must be in debug, if debug enabled (1)`)
	}

	TestsConfigureAndInitialize()

	if Debug && !IsDebugEnabled() {
		t.Fatal(`Must not be in debug, if debug enabled (2)`)
	}
}

func TestIsEnvDevelopment(t *testing.T) {
	if IsEnvDevelopment() {
		t.Fatal(`Must not be in development`)
	}

	AppEnvironment = APP_ENVIRONMENT_DEVELOPMENT

	if !IsEnvDevelopment() {
		t.Fatal(`Must be in development`)
	}

	TestsConfigureAndInitialize()

	if IsEnvDevelopment() {
		t.Fatal(`Must not be in development`)
	}
}

func TestIsEnvProduction(t *testing.T) {
	if IsEnvProduction() {
		t.Fatal(`Must not be in production`)
	}

	AppEnvironment = APP_ENVIRONMENT_PRODUCTION

	if !IsEnvProduction() {
		t.Fatal(`Must be in production`)
	}

	TestsConfigureAndInitialize()

	if IsEnvProduction() {
		t.Fatal(`Must not be in production`)
	}
}

func TestIsEnvLocal(t *testing.T) {
	if IsEnvLocal() {
		t.Fatal(`Must not be in local`)
	}

	AppEnvironment = APP_ENVIRONMENT_LOCAL

	if !IsEnvLocal() {
		t.Fatal(`Must be in local`)
	}

	TestsConfigureAndInitialize()

	if IsEnvLocal() {
		t.Fatal(`Must not be in testing`)
	}
}

func TestIsEnvTesting(t *testing.T) {
	AppEnvironment = APP_ENVIRONMENT_LOCAL // reset to local, as gets set to testing during tests

	if IsEnvTesting() {
		t.Fatal(`Must not be in testing`)
	}

	TestsConfigureAndInitialize()

	if !IsEnvTesting() {
		t.Fatal(`Must not be in testing`)
	}
}
