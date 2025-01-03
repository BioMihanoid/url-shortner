package app

import (
	"github.com/BioMihanoid/url-shortner/internal/config"
	"github.com/BioMihanoid/url-shortner/internal/handler"
	"github.com/BioMihanoid/url-shortner/internal/repository"
	"github.com/BioMihanoid/url-shortner/internal/service"
	"github.com/BioMihanoid/url-shortner/internal/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Start() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(utils.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
		logrus.Fatalf("error starting server: %s", err.Error())
	}
}
