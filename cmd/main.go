package main

import (
	"log"

	todo "github.com/kahuri1/message-processor"
	"github.com/kahuri1/message-processor/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
