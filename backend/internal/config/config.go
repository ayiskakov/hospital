package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"serviceName"`
}

type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Database struct {
	Dsn string `yaml:"dsn"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project  Project  `yaml:"project"`
	Rest     Rest     `yaml:"rest"`
	Database Database `yaml:"database"`
}

var cfg *Config

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(configYML string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}

func Get() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}
