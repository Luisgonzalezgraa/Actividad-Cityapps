package main

import (
	"context"
	"log"
	"time"

	"citiaps-tasks-backend/internal/config"
	"citiaps-tasks-backend/internal/db"
	"citiaps-tasks-backend/internal/handlers"
	"citiaps-tasks-backend/internal/repositories"
	"citiaps-tasks-backend/internal/routes"
	"citiaps-tasks-backend/internal/services"
)

func main() {
	cfg := config.Load()

	client, collection, err := db.Connect(cfg)
	if err != nil {
		log.Fatal("error conectando a MongoDB: ", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = client.Disconnect(ctx)
	}()

	repo := repositories.NewTaskRepository(collection)
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)
	router := routes.SetupRouter(handler)

	log.Printf("backend corriendo en puerto %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("error levantando servidor: ", err)
	}
}
