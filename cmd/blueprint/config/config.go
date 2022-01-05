package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type appConfig struct {
	//Example variable
	ConfigVar string
}

var Config appConfig

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigFile("example") // filename config
	v.SetConfigFile("yaml")
	v.SetEnvPrefix("blueprint")
	v.AutomaticEnv()

	for _, path := range configPaths {
		v.AddConfigPath(path) // path to look for the config file in
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	return v.Unmarshal(&Config)

}
