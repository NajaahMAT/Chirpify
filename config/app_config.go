package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration.
type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

// LoadConfig loads configuration from a YAML file.
func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
