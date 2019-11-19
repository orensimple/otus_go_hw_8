package config

import (
	"log"

	"github.com/spf13/viper"
)

// Init config
func Init(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found %s", err.Error())
		} else {
			log.Fatalf("Config file was found but another error was produced %s", err.Error())
		}
	}
	Validate()

	return viper.ReadInConfig()
}

// Validate check config params
func Validate() {
	if len(viper.GetString("log_level.file")) <= 0 {
		log.Fatalf("Cannot read log_level.file in config")
	}
	if len(viper.GetString("log_level.command")) <= 0 {
		log.Fatalf("Cannot read log_level.ficommandle in config")
	}
	if len(viper.GetString("log_file")) <= 0 {
		log.Fatalf("Cannot read log_file in config")
	}
	if len(viper.GetString("http_listen.ip")) <= 0 {
		log.Fatalf("Cannot read http_listen.ip in config")
	}
	if len(viper.GetString("http_listen.port")) <= 0 {
		log.Fatalf("Cannot read lhttp_listen.port in config")
	}
}
