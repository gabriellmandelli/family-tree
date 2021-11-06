package config

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var config Configuration

//Configuration struct
type Configuration struct {
	AppName    string
	AppPort    string
	HealthPort string
}

func initConfig() *viper.Viper {
	viperConfig := viper.New()

	viperConfig.SetConfigName("configuration")
	viperConfig.SetConfigType("yaml")
	viperConfig.AddConfigPath("internal/config")

	setDefaults(viperConfig)

	if err := viperConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
		}
	}

	return viperConfig
}

func setDefaults(viperConfig *viper.Viper) {
	viperConfig.SetDefault("app.name", "family-tree")
	viperConfig.SetDefault("app.port", ":8080")
	viperConfig.SetDefault("health.port", ":8081")
	viperConfig.SetDefault("health.port", ":8081")
}

//GetConfig return Configuration
func GetConfig() Configuration {
	once.Do(func() {
		viperConfig := initConfig()
		config = Configuration{
			AppName:    viperConfig.GetString("app.name"),
			AppPort:    viperConfig.GetString("app.port"),
			HealthPort: viperConfig.GetString("health.port"),
		}
	})
	return config
}
