package main

import (
	"os"

	"github.com/joho/godotenv"
	todo "github.com/kahuri1/message-processor"
	"github.com/kahuri1/message-processor/pkg/handler"
	"github.com/kahuri1/message-processor/pkg/repository"
	"github.com/kahuri1/message-processor/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error load env: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Errorf("Failed initialization db: %s", err.Error())
		//logrus.Fatalf("Failed initialization db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos, nil)
	handlers := handler.Newhandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("C:/proect/message-processor/configs") //исправить на нормальный путь
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
