package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfiguration() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("could not load config %w", err))
	}
}