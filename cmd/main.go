package main

import (
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/handler"
	"github.com/oneils/todo-app/pkg/repository"
	"github.com/oneils/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.New()
	services := service.New(repos)
	handlers := handler.New(services)

	server := new(todo.Server)

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while starting http server %s", err.Error())
	}
}
