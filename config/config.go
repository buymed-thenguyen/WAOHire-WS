package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Port     string `yaml:"port"`
	GrpcPort string `yaml:"grpc_port"`
}

const _PATH = "./config.yml"

func Load() (*Config, error) {
	data, err := os.ReadFile(_PATH)
	if err != nil {
		log.Printf("read config file: %w", err)
		return nil, err
	}

	var appConfig *Config
	if err = yaml.Unmarshal(data, &appConfig); err != nil {
		log.Printf("unmarshal yaml: %w", err)
		return nil, err
	}

	if appConfig == nil {
		log.Printf("config nil")
		return nil, errors.New("invalid config")
	}

	return appConfig, nil
}
