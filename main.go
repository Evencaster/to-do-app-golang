package main

import (
	"github.com/Evencaster/to-do-app-golang/controllers"
	"github.com/Evencaster/to-do-app-golang/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/webapp", "./frontend/build")
	repo := repository.NewMemRepo()
	taskController := controllers.NewTaskController(repo)
	r.GET("/", taskController.GetAllTasks)
	r.POST("/", taskController.AddTask)
	r.DELETE("/:id", taskController.RemoveTask)
	r.DELETE("/", taskController.RemoveAllTasks)
	err := r.Run("127.0.0.1:4444")
	if err != nil {
		panic(err)
	}
}
