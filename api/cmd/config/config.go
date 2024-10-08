package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type postgres struct {
	Username string
	Password string
	Host     string
	Port     int
}

type Config struct {
	Port      int
	SecretKey string
	Postgres  postgres
}

var Cfg *Config = &Config{}

func LoadConfig() error {
	contents, err := os.ReadFile("config.toml")

	if err != nil {
		return err
	}

	err = toml.Unmarshal(contents, Cfg)

	if err != nil {
		return err
	}

	return nil
}
