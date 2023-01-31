package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() error {
	logrus.Info("[config] Init...")
	var err error = nil
	config := ConfigType{}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	MongoUri = config.MongoUri
	Server = config.Server

	logrus.Info("[config] Init Success")
	return nil
}
