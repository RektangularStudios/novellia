package main

import (
	"fmt"
	"os"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	CardanoGraphQL struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"cardano-graphql"`
	Postgres struct {
		Database string `yaml:"database"`
		Host string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Mocked bool
}

func getConfigPath() (string, error) {
	if (len(os.Args) != 2) {
		return "", fmt.Errorf("expected config path as first argument")
	}
	
	configPath := os.Args[1]
	return configPath, nil
}

func loadConfig(configPath string) (*Config, error) {
	var config Config

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
