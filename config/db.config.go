package config

import (
	"log"
	"os"
)

type DatabaseConfig struct {
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	Collection string
}

func GetDbConfig() DatabaseConfig {
	dbName := os.Getenv("DB")
	if dbName == "" {
		log.Fatalf("Failed to load environment variable Database")
	}

	dbUser := os.Getenv("USER")
	if dbUser == "" {
		log.Fatalf("Failed to load environment variable USER")
	}

	dbPass := os.Getenv("PASS")
	if dbPass == "" {
		log.Fatalf("Failed to load environment variable PASS")
	}

	colName := os.Getenv("COLLECTION")
	if colName == "" {
		log.Fatalf("Failed to load environment variable COLLECTION")
	}

	dbPort := os.Getenv("DBPORT")
	if dbPort == "" {
		log.Fatalf("Failed to load environment variable DBPORT")
	}

	return DatabaseConfig{
		DBPort:     dbPort,
		DBUsername: dbUser,
		DBPassword: dbPass,
		DBName:     dbName,
		Collection: colName,
	}
}
