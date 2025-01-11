package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB DB `yaml:"db"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

func NewFromFile() (*Config, error) {
	configPath := "configs/config.yaml"

	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
