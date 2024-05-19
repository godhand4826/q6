package config

import (
	"cmp"
	"os"
	"q6/lib/config"
)

type Config struct {
	Addr string
}

func NewConfig() *Config {
	config.LoadToEnv()

	return &Config{
		Addr: cmp.Or(os.Getenv("ADDR"), ":8080"),
	}
}
