package main

import (
	"ImGoDeveloper"
	"ImGoDeveloper/pkg/handler"
	"ImGoDeveloper/pkg/repository"
	"ImGoDeveloper/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initialization config %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "gipernova",
		Password: "qwerty",
		DBname:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	servises := service.NewService(repos)
	handlers := handler.NewHandler(servises)

	srv := new(ImGoDeveloper.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
