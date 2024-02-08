package main

import (
	"github.com/Blxssy/todo-app"
	"github.com/Blxssy/todo-app/pkg/handler"
	"github.com/Blxssy/todo-app/pkg/repository"
	"github.com/Blxssy/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
