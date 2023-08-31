package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"internship_avito"
	"internship_avito/pkg/handler"
	"internship_avito/pkg/repository"
	"internship_avito/pkg/services"
)

// @Title Internship-avito
// @Version 1.0
// @Description API Service for Internship-avito, user and segment

// @localhost:8000
// @BasePath
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config, err: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLmode:  "disable",

		// Локальная версия бд
		//Host:     "localhost",
		//Port:     "5433",
		//Username: "postgres",
		//Password: "1234",
		//DBName:   "postgres",
		//SSLmode:  "disable",
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
