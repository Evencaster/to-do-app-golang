package repository

type Repo interface {
	GetAllTasks() []Task
	AddTask(name string) uint64
	RemoveTask(id uint64)
	RemoveAllTasks()
}
