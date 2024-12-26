package app

import (
	"github.com/BioMihanoid/url-shortner/internal/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Start() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}
