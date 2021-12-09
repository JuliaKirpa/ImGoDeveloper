package main

import (
	I_mGoDeveloper "I-mGoDeveloper"
	"log"
)

func main(){
	srv := new(I_mGoDeveloper.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
