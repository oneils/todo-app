package main

import (
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/handler"
	"github.com/oneils/todo-app/pkg/repository"
	"github.com/oneils/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err.Error())
	}

	repos := repository.New()
	services := service.New(repos)
	handlers := handler.New(services)

	server := new(todo.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while starting http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
