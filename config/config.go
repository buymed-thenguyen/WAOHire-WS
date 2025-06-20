package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Port     string  `yaml:"port"`
	GrpcPort string  `yaml:"grpc_port"`
	Logger   *Logger `yaml:"logger"`
}

type Logger struct {
	Path string `json:"path"`
}

const _FILE = "config.yml"

func Load() (*Config, error) {
	env := os.Getenv("ENV")
	path := "./config/local/"
	if env == "prd" {
		path = "./config/prd/"
	}

	data, err := os.ReadFile(path + _FILE)
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

	if env == "prd" {
		port := os.Getenv("PORT") // fallback for railway deployment
		if port == "" {
			port = appConfig.Port
		}
	}

	return appConfig, nil
}
