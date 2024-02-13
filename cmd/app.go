package main

import (
	"AuthService/config"
	lg "AuthService/pkg/logger"
	"github.com/spf13/viper"
)

func main() {

	logger := lg.InitLogger()

	logger.Info.Print("Executing InitConfig.")
	if err := config.InitConfig(); err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Println("SERVER PORT: " + viper.GetString("port"))

	logger.Info.Println("Connecting to MongoDB.")

}
