package main

import (
	"ImGoDeveloper"
	"ImGoDeveloper/pkg/handler"
	"ImGoDeveloper/pkg/repository"
	"ImGoDeveloper/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {

	repos := repository.NewRepository()
	servises := service.NewService(repos)
	handlers := handler.NewHandler(servises)

	srv := new(ImGoDeveloper.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}

}
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
