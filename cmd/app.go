package main

import (
	"AuthService/config"
	"AuthService/internal/repository"
	"AuthService/internal/server"
	"AuthService/internal/service"
	lg "AuthService/pkg/logger"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := lg.InitLogger()

	logger.Info.Print("Executing InitConfig.")
	if err := config.InitConfig(); err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Println("Connecting to MongoDB.")
	client, err := repository.NewMongoClient("mongodb",
		viper.GetString("MONGO_HOST"),
		viper.GetString("MONGO_PORT"),
		viper.GetString("MONGO_INITDB_ROOT_USERNAME"),
		viper.GetString("MONGO_INITDB_ROOT_PASSWORD"))

	if err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Println(fmt.Sprintf("Created collection '%s' in database '%s'", repository.CollectionName, repository.DatabaseName))
	if err := repository.CreateCollections(client); err != nil {
		logger.Err.Fatalf(err.Error())
	}

	logger.Info.Println("Executing NewRepository")
	repo := repository.NewRepository(client, logger)

	logger.Info.Println("Executing NewServices")
	services := service.NewServices(repo, logger)

	port := viper.GetString("port")
	logger.Info.Println("Executing NewHttpServer")
	srv := server.NewHttpServer(services, logger, ":"+port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Err.Fatalf("error occured while running http server: \"%s\" \n", err.Error())
		}
	}()

	logger.Info.Print("Starting the server on port: " + port + "\n\n")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info.Println("Server Shutting Down.")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Err.Fatalf("error occured while server shutting down: \"%s\" \n", err.Error())
	}

	logger.Info.Println("DB connection closing.")

	if err := client.Disconnect(context.TODO()); err != nil {
		logger.Err.Fatalf("error occured on db connection close: \"%s\" \n", err.Error())
	}
}
