package main

import (
	"log"

	"github.com/kibah52724/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
