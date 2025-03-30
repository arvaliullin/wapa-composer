package app

import (
	"os"

	"gopkg.in/yaml.v3"
)

// ComposerConfig настройки сервиса
type ComposerConfig struct {
	DbConnection string `yaml:"database_connection"`
	Address      string `yaml:"composer_address"`
}

func NewComposerConfig(configPath string) (*ComposerConfig, error) {
	config := &ComposerConfig{}
	err := config.Load(configPath)
	return config, err
}

func (config *ComposerConfig) Load(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}
	return config.overrideEnv()
}

func (config *ComposerConfig) overrideEnv() error {
	if dbConn, exists := os.LookupEnv("COMPOSER_DB_CONNECTION"); exists {
		config.DbConnection = dbConn
	}
	return nil
}
