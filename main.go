package main

import (
	"github.com/Evencaster/to-do-app-golang/controllers"
	"github.com/Evencaster/to-do-app-golang/repository/repos/mem-repo"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo := mem_repo.NewMemRepo()
	taskController := controllers.NewTaskController(repo)

	r.GET("/tasks", taskController.GetAllTasks)
	r.POST("/tasks", taskController.AddTask)
	r.DELETE("/tasks/:id", taskController.RemoveTask)
	r.DELETE("/tasks", taskController.RemoveAllTasks)

	r.Use(static.Serve("/", static.LocalFile("./frontend/build", false)))

	err := r.Run("127.0.0.1:5555")
	if err != nil {
		panic(err)
	}
}
