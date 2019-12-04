package config_test

import (
	"os"
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
