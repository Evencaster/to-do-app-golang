package controllers

import (
	"github.com/Evencaster/to-do-app-golang/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddTaskBody struct {
	Name string `json:"name"`
}

type TaskController struct {
	repo repo.Repo
}

func NewTaskController(repo repo.Repo) *TaskController {
	return &TaskController{repo: repo}
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks := c.repo.GetAllTasks()
	ctx.JSON(200, tasks)
}

func (c *TaskController) AddTask(ctx *gin.Context) {
	var task AddTaskBody
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "bad json",
		})
	}

	id := c.repo.AddTask(task.Name)

	ctx.JSON(200, gin.H{
		"ID": strconv.Itoa(int(id)),
	})
}

func (c *TaskController) RemoveTask(ctx *gin.Context) {
	idStr, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "bad id param",
		})
	}

	c.repo.RemoveTask(uint64(id))
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func (c *TaskController) RemoveAllTasks(ctx *gin.Context) {
	c.repo.RemoveAllTasks()
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
