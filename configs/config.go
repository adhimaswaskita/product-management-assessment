package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// YAMLConfigLoader is the loader of YAML file configuration.
type YAMLConfigLoader struct {
	fileLocation string
}

// NewYamlConfigLoader return the YAML Configuration loader.
func NewYamlConfigLoader(fileLocation string) *YAMLConfigLoader {
	return &YAMLConfigLoader{
		fileLocation: fileLocation,
	}
}

// ServiceConfig stores the whole configuration for service.
type ServiceConfig struct {
	ServiceData ServiceDataConfig `yaml:"service_data"`
	SourceData  SourceDataConfig  `yaml:"source_data"`
}

// ServiceDataConfig contains the service data configuration.
type ServiceDataConfig struct {
	Address string `yaml:"address"`
}

// SourceDataConfig contains the source data configuration.
type SourceDataConfig struct {
	DBServer   string `yaml:"db_server"`
	DBPort     int    `yaml:"db_port"`
	DBName     string `yaml:"db_name"`
	DBUsername string `yaml:"db_username"`
	DBPassword string `yaml:"db_password"`
	DBTimeout  int    `yaml:"db_timeout"`
}

func getRawConfig(fileLocation string) (*ServiceConfig, error) {
	configByte, err := os.ReadFile(fileLocation)
	if err != nil {
		return nil, err
	}
	config := &ServiceConfig{}
	err = yaml.Unmarshal(configByte, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// GetServiceConfig parse the configuration from YAML file.
func (c *YAMLConfigLoader) GetServiceConfig() (*ServiceConfig, error) {
	config, err := getRawConfig(c.fileLocation)
	if err != nil {
		return nil, fmt.Errorf("unable to get raw config content: %v", err)
	}
	return config, nil
}
