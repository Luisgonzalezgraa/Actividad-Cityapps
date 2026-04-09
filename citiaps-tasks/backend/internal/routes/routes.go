package routes

import (
	"citiaps-tasks-backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskHandler *handlers.TaskHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	_ = r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetTasks)
	r.PUT("/tasks/:id/complete", taskHandler.CompleteTask)
	r.PUT("/tasks/:id/recover", taskHandler.RecoverTask)
	r.PUT("/tasks/:id/toggle-active", taskHandler.ToggleActive)
	r.DELETE("/tasks/:id/permanent", taskHandler.DeleteTaskPermanent)
	r.GET("/tasks/:id", taskHandler.GetTaskByID)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	return r
}
