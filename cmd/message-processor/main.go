package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	todo "github.com/kahuri1/message-processor"
	"github.com/kahuri1/message-processor/pkg/handler"
	"github.com/kahuri1/message-processor/pkg/repository"
	"github.com/kahuri1/message-processor/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load env: %s", err.Error())
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
		log.Fatalf("Failed initialization db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.Newhandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("C:/proect/message-processor/configs") //исправить на нормальный путь
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
