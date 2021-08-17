package repo

import "github.com/Evencaster/to-do-app-golang/entities"

type Repo interface {
	GetAllTasks() []entities.Task
	AddTask(name string) uint64
	RemoveTask(id uint64)
	RemoveAllTasks()
}
