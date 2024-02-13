package config

import "github.com/spf13/viper"

func InitConfig() error {

	viper.SetConfigFile("config/.env")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	viper.SetConfigFile("config/config.yml")
	err = viper.MergeInConfig()

	return err
}
