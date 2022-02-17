package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	PortHTTP      string `yaml:"port"`
	UrlBackend    string `yaml:"backend"`
	CapacityCache int    `yaml:"capacity"`
}

func NewConfig(cfgPath string) (*Config, error) {
	var cfg Config
	file, err := os.Open(cfgPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	d := yaml.NewDecoder(file)
	if err := d.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
