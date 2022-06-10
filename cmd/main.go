package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"todo-app/pkg/handler"
	"todo-app/pkg/repo"
	"todo-app/pkg/service"

	todo "todo-app"
)

func main() {
	err := initConfig()
	if err != nil {
		logrus.Fatalf("error initialisation configs: %s", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}

	db, err := repo.NewDB(repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	repo := repo.New(db)
	service := service.New(repo)
	handler := handler.New(service)

	srv := new(todo.Server)
	err = srv.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
