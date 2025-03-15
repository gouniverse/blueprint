package config_v2

import (
	"os"
	"testing"
)

func TestInitialize_AppServerHostAndPort(t *testing.T) {
	cfg, err := TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.WebServerHost == "" {
		t.Fatal("WebServerHost SHOULD NOT BE empty")
	}

	if cfg.WebServerPort == "" {
		t.Fatal("WebServerPort SHOULD NOT BE empty")
	}

	if cfg.AppUrl == "" {
		t.Fatal("AppUrl SHOULD NOT BE empty")
	}

	if cfg.DbDriver == "" {
		t.Fatal("DbDriver SHOULD NOT BE empty")
	}

	if cfg.DbHost != "" {
		t.Fatal("DbHost SHOULD BE empty")
	}

	if cfg.DbName == "" {
		t.Fatal("DbName SHOULD NOT BE empty")
	}

	if cfg.DbPass != "" {
		t.Fatal("DbPass SHOULD BE empty")
	}

	if cfg.DbPort != "" {
		t.Fatal("DbPort SHOULD BE empty")
	}
}

func TestInitialize_Debug(t *testing.T) {
	os.Setenv("DEBUG", "yes")
	cfg, err := TestsConfigureAndInitialize()

	if err != nil {
		t.Fatal(err)
	}

	if cfg.Debug == false {
		t.Fatal("Debug SHOULD NOT BE false")
	}
	if cfg.WebServerHost == "" {
		t.Fatal("ServerHost SHOULD NOT BE empty")
	}
	if cfg.WebServerPort == "" {
		t.Fatal("ServerPort SHOULD NOT BE empty")
	}
}
