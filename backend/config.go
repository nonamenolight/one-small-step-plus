package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	SiteName  string `yaml:"site_name" json:"site_name"`
	StartDate string `yaml:"start_date" json:"start_date"`
}

func loadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
