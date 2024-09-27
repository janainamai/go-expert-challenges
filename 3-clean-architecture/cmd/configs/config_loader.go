package configs

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func LoadConfigs() *Config {
	scope := getScope()

	filePath := fmt.Sprintf("cmd/configs/%s.yml", scope)
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("error reading file config %s: %s", scope, err.Error()))
	}

	var config *Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(fmt.Sprintf("error unmarshal file content %s: %s", scope, err.Error()))
	}

	return config
}

func getScope() string {
	scope := os.Getenv("SCOPE")

	if scope == "" {
		scope = "local"
	}

	logrus.Infof("Scope: %s", scope)
	return scope
}
