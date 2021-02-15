package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/oneils/todo-app"
	"github.com/oneils/todo-app/pkg/handler"
	"github.com/oneils/todo-app/pkg/repository"
	"github.com/oneils/todo-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
	"os/signal"
	"syscall"
)

// @title TodoList APP
// @version 1.0
// @description API Service for TodoList application

// @host localhost:8080
// @basePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while reading config file %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("Error while loading env file %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslMode"),
	})
	if err != nil {
		logrus.Fatalf("Error while connecting to DB. Error: %s", err.Error())
	}

	repos := repository.New(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)

	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error while starting http server %s", err.Error())
		}
	}()

	logrus.Printf("TODO App started on http://localhost:%s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("TODO App is shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred while shutting down the server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred while closing DB connectionsr: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
