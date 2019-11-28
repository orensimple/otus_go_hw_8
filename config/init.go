package config

import (
	"github.com/golang/glog"

	"github.com/spf13/viper"
)

// Init config
func Init(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			glog.Errorf("Config file not found %s", err.Error())
		} else {
			glog.Errorf("Config file was found but another error was produced %s", err.Error())
		}
	}
	Validate()

	return viper.ReadInConfig()
}

// Validate check config params
func Validate() {
	if len(viper.GetString("log_level.file")) <= 0 {
		glog.Errorf("Cannot read log_level.file in config")
	}
	if len(viper.GetString("log_level.command")) <= 0 {
		glog.Errorf("Cannot read log_level.ficommandle in config")
	}
	if len(viper.GetString("log_file")) <= 0 {
		glog.Errorf("Cannot read log_file in config")
	}
	if len(viper.GetString("http_listen.ip")) <= 0 {
		glog.Errorf("Cannot read http_listen.ip in config")
	}
	if len(viper.GetString("http_listen.port")) <= 0 {
		glog.Errorf("Cannot read lhttp_listen.port in config")
	}
}
