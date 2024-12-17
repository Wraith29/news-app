package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type postgres struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

type Config struct {
	SecretKey string   `toml:"secret-key"`
	Postgres  postgres `toml:"postgres"`
}

var Cfg *Config = &Config{}

func Load() error {
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
