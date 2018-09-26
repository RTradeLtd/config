package config_test

import (
	"os"
	"testing"

	"github.com/RTradeLtd/Temporal/config"
)

func TestGenerateAndLoadConfig(t *testing.T) {
	testconf := "./testconfig.json"
	defer os.Remove(testconf)
	if err := config.GenerateConfig(testconf); err != nil {
		t.Fatal(err)
	}
	if _, err := config.LoadConfig(testconf); err != nil {
		t.Fatal(err)
	}
}
