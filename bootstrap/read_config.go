package bootstrap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
)

type ConfigOptions struct {
	BindAddress        string `json:"bind_address"`
	DBConnectionString string `json:"db_connection_string"`
	JWTSigningKey      string `json:"jwt_signing_key"`
}

func ReadConfig(configFilePath string, useTLS bool) (bool, *ConfigOptions) {
	fi, err := os.Stat(configFilePath)
	if err != nil {
		panic(err)
	}

	if fi.IsDir() {
		panic(fmt.Errorf("config file is a directory"))
	}

	b, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Errorf("error reading configuration file: %s", err.Error()))
	}

	configOptions := &ConfigOptions{}
	err = json.Unmarshal(b, configOptions)
	if err != nil {
		panic(fmt.Errorf("error parsing json configuration file: %s", err.Error()))
	}

	logrus.
		WithField("config_file", configFilePath).
		WithField("config", configOptions).
		Info("Configuration Loaded")

	return useTLS, configOptions
}
