package main

import (
	"github.com/EgorBKV4/todo-app"
	"github.com/EgorBKV4/todo-app/pkg/handler"
	"github.com/EgorBKV4/todo-app/pkg/repository"
	"github.com/EgorBKV4/todo-app/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	DB, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("DB.host"),
		Port:     viper.GetString("DB.port"),
		Username: viper.GetString("DB.username"),
		DBName:   viper.GetString("DB.dbname"),
		SSLMode:  viper.GetString("DB.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize DB: %s", err.Error())
	}

	repos := repository.NewRepository(DB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
