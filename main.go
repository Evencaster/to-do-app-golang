package main

import (
	"github.com/Evencaster/to-do-app-golang/controllers"
	mysql_repo "github.com/Evencaster/to-do-app-golang/repository/repos/mysql-repo"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	//repo := mem_repo.NewMemRepo()
	repo := mysql_repo.NewMySQLRepo()
	taskController := controllers.NewTaskController(repo)

	r.GET("/tasks", taskController.GetAllTasks)
	r.POST("/tasks", taskController.AddTask)
	r.DELETE("/tasks/:id", taskController.RemoveTask)
	r.DELETE("/tasks", taskController.RemoveAllTasks)

	r.Use(static.Serve("/", static.LocalFile("./frontend/build", false)))

	err = r.Run("127.0.0.1:5555")
	if err != nil {
		panic(err)
	}
}
