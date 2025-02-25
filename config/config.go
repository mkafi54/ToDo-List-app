package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

func SetupConfig() error {
	var configuration *Configuration

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error to reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("error to decode, %v", err)
		return err
	}

	return nil
}
