package config

import "github.com/spf13/viper"

func NewViperConfig() error {
	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
