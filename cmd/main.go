package main

import (
	"log"

	todo "github.com/kahuri1/message-processor"
	"github.com/kahuri1/message-processor/pkg/handler"
	"github.com/kahuri1/message-processor/pkg/repository"
	"github.com/kahuri1/message-processor/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.Newhandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
