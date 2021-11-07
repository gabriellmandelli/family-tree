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
	MongoDB    database
}

type database struct {
	URI      string
	UserName string
	Password string
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
	viperConfig.SetDefault("mongodb.uri", "mongodb://localhost:27017")
	viperConfig.SetDefault("mongodb.username", "test")
	viperConfig.SetDefault("mongodb.password", "test")
}

//GetConfig return Configuration
func GetConfig() Configuration {
	once.Do(func() {
		viperConfig := initConfig()
		config = Configuration{
			AppName:    viperConfig.GetString("app.name"),
			AppPort:    viperConfig.GetString("app.port"),
			HealthPort: viperConfig.GetString("health.port"),
			MongoDB: database{
				URI:      viperConfig.GetString("mongodb.uri"),
				UserName: viperConfig.GetString("mongodb.username"),
				Password: viperConfig.GetString("mongodb.password"),
			},
		}
	})
	return config
}
