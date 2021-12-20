package main

import (
	"ImGoDeveloper"
	"ImGoDeveloper/pkg/handler"
	"ImGoDeveloper/pkg/repository"
	"ImGoDeveloper/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initialization config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error load env variable: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	servises := service.NewService(repos)
	handlers := handler.NewHandler(servises)

	srv := new(ImGoDeveloper.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
