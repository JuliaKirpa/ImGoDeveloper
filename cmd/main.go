package main

import (
	"ImGoDeveloper"
	"log"
)

func main(){
	srv := new(ImGoDeveloper.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
