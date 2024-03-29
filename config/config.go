package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token string `yaml:"token"`
}

func GetToken() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "config.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg.Token
}
