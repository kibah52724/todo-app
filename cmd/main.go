package main

import (
	"log"

	"github.com/kibah52724/todo-app"
	"github.com/kibah52724/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
