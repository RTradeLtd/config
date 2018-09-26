package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// LoadConfig loads a TemporalConfig from given filepath
func LoadConfig(configPath string) (*TemporalConfig, error) {
	var tCfg TemporalConfig
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &tCfg)
	if err != nil {
		return nil, err
	}
	return &tCfg, nil
}

// GenerateConfig writes a empty TemporalConfig template to given filepath
func GenerateConfig(configPath string) error {
	template := &TemporalConfig{}
	b, err := json.Marshal(template)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, b, os.ModePerm)
}
