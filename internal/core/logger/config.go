package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LOG_LEVEL" required:"true"`
	Folder string `envconfig:"LOG_FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(fmt.Errorf("get logger config: %w", err))
	}

	return cfg
}
