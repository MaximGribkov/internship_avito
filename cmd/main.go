package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"internship_avito"
	"internship_avito/pkg/handler"
	"internship_avito/pkg/repository"
	"internship_avito/pkg/services"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config, err: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "1234",
		Password: "1234",
		DBName:   "1234",
		SSLmode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db, err: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := services.NewService(repos)
	handlers := handler.NewHandler(service)

	svr := new(internship_avito.Server)
	if err := svr.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error start server, err: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
