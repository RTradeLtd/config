package config_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/RTradeLtd/config/v2"
)

func TestGenerateAndLoadConfig(t *testing.T) {
	testconf := "./example.config.json"
	defer os.Remove(testconf)
	if err := config.GenerateConfig(testconf); err != nil {
		t.Fatal(err)
	}
	if _, err := config.LoadConfig(testconf); err != nil {
		t.Fatal(err)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("API_CONNECTION_LISTENADDRESS", "0.0.0.0:9090")
	conf, err := config.LoadConfig("")
	if err != nil {
		t.Fatal(err)
	}
	if conf.API.Connection.ListenAddress != "0.0.0.0:9090" {
		t.Fatal("bad api listen address")
	}
	if reflect.DeepEqual(&config.TemporalConfig{}, conf) {
		t.Fatal("should not be equal")
	}
	os.Unsetenv("API_CONNECTION_LISTENADDRESS")
	conf, err = config.LoadConfig("")
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestGenerateConfigFailure(t *testing.T) {
	testConf := "/root/toor/config.json"
	if err := config.GenerateConfig(testConf); err == nil {
		t.Fatal("error expected")
	}
}

func TestLoadConfigFailure(t *testing.T) {
	testFileExists := "./README.md"
	if _, err := config.LoadConfig(testFileExists); err == nil {
		t.Fatal("error expected")
	}
	testFileNotExists := "/root/toor/config.json"
	if _, err := config.LoadConfig(testFileNotExists); err == nil {
		t.Fatal("error expected")
	}
}
