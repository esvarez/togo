package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	App
}

var (
	home, _    = os.UserHomeDir()
	finitoDir  = home + "/.togo/"
	configFile = finitoDir + "config.yaml"
)

type App struct {
	DefaultList string `mapstructure:"default_list" yaml:"default_list"`
}

func LoadConfig() (*Configuration, error) {
	file := viper.New()
	file.SetConfigFile(configFile)

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

func SaveConfiguration(cfg *Configuration) error {
	file := viper.New()
	file.SetConfigFile(configFile)
	file.SetConfigType("yaml")
	file.Set("app", cfg.App)

	if err := file.WriteConfig(); err != nil {
		return err
	}

	return nil
}

func init() {
	file := viper.New()
	file.SetConfigFile(configFile)
	file.AddConfigPath(finitoDir)
	file.SetConfigType("yaml")
	file.Set("app.sheet_id", nil)
	file.SafeWriteConfig()
}
