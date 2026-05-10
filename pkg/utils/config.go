package utils

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Timezones []string `yaml:"timezones"`
}

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".tmz.yaml"), nil
}

func LoadConfig() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	f, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Config{Timezones: []string{}}, nil
		}
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}

	if config.Timezones == nil {
		config.Timezones = []string{}
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
