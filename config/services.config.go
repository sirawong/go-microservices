package config

import (
	"log"
	"os"
)

type ServiceConfig struct {
	AppPort string
}

func GetServiceConfig() ServiceConfig {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatalf("Failed to load environment variable appPort")
	}

	return ServiceConfig{
		AppPort: appPort,
	}
}
