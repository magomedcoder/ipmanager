package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Postgres *Postgres `yaml:"postgres"`
	Jwt      *Jwt      `yaml:"jwt"`
}

type Jwt struct {
	Secret      string `yaml:"secret"`
	ExpiresTime int64  `yaml:"expires_time"`
}

func New(filename string) *Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var conf Config
	if yaml.Unmarshal(content, &conf) != nil {
		panic(fmt.Sprintf("%s: %v", filename, err))
	}

	return &conf
}
