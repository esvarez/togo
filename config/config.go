package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Configuration struct {
	App
}

type App struct {
	DefaultList string `mapstructure:"default_list" yaml:"default_list"`
}

const path = "./config/config.yaml"

func NewConfig() (*Configuration, error) {
	file := viper.New()
	file.SetConfigFile(path)

	if err := file.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Configuration{}
	if err := file.Unmarshal(cfg); err != nil {
		return nil, err
	}

	v := validator.New()
	if err := v.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
