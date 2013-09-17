package helpers

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

var configFile *yaml.File

func Config() *yaml.File {
	if configFile != nil {
		return configFile
	}
	file := "config.yaml"
	yamlConfig, _ := yaml.ReadFile(file)
	configFile = yamlConfig
	return configFile
}
