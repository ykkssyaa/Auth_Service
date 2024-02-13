package main

import (
	"AuthService/config"
	"AuthService/internal/repository"
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
	_, err := repository.NewMongoClient("mongodb",
		viper.GetString("MONGO_HOST"),
		viper.GetString("MONGO_PORT"),
		viper.GetString("MONGO_INITDB_ROOT_USERNAME"),
		viper.GetString("MONGO_INITDB_ROOT_PASSWORD"))

	if err != nil {
		logger.Err.Fatalf(err.Error())
	}

}
