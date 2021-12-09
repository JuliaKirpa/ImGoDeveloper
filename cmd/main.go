package main

import (
	"ImGoDeveloper"
	"ImGoDeveloper/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(ImGoDeveloper.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}

}
