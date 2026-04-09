package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	MongoURI   string
	Database   string
	Collection string
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		Port:       getEnv("PORT", "8080"),
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Database:   getEnv("MONGO_DB", "tasksdb"),
		Collection: getEnv("MONGO_COLLECTION", "tasks"),
	}

	if cfg.MongoURI == "" {
		log.Fatal("MONGO_URI no configurado")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
