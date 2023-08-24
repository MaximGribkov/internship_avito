package main

import (
	"github.com/sirupsen/logrus"
	"internship_avito"
	"internship_avito/pkg/handler"
	"internship_avito/pkg/repository"
	"internship_avito/pkg/services"
)

func main() {
	repos := repository.NewRepository()
	service := services.NewService(repos)
	handlers := handler.NewHandler(service)

	svr := new(internship_avito.Server)
	if err := svr.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error start server, err: %s", err.Error())
	}
}
