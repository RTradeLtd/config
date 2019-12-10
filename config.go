package config

import (
	"bytes"
	"encoding/json"
	"github.com/vrischmann/envconfig"
	"io/ioutil"
	"os"
	"reflect"
)

// LoadConfig loads a TemporalConfig from given filepath
func LoadConfig(configPath string) (*TemporalConfig, error) {
	// if configPath is empty, load from env
	if configPath == "" {
		conf, err := LoadConfigFromEnv()
		if err != nil {
			return nil, err
		}
		// this will pass if we failed to pull config
		// from the environment, and will then default
		// to config file path based loading
		if !reflect.DeepEqual(&TemporalConfig{}, conf) {
			return conf, nil
		}
	}
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var tCfg TemporalConfig
	if err = json.Unmarshal(raw, &tCfg); err != nil {
		return nil, err
	}

	tCfg.setDefaults()

	return &tCfg, nil
}

// LoadConfigFromEnv is used to load a TemporalConfig object us env vars
func LoadConfigFromEnv() (*TemporalConfig, error) {
	cfg := &TemporalConfig{}
	err := envconfig.Init(cfg)
	return cfg, err
}

// GenerateConfig writes a empty TemporalConfig template to given filepath
func GenerateConfig(configPath string) error {
	template := &TemporalConfig{}
	template.setDefaults()
	b, err := json.Marshal(template)
	if err != nil {
		return err
	}

	var pretty bytes.Buffer
	if err = json.Indent(&pretty, b, "", "\t"); err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, pretty.Bytes(), os.ModePerm)
}

func (t *TemporalConfig) setDefaults() {
	if t.LogDir == "" {
		t.LogDir = "/var/log/temporal/"
	}
	if len(t.API.Connection.CORS.AllowedOrigins) == 0 {
		t.API.Connection.CORS.AllowedOrigins = []string{"temporal.cloud", "backup.temporal.cloud"}
	}
}
