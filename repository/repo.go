package repo

type Task struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}

type Repo interface {
	GetAllTasks() []Task
	AddTask(name string) uint64
	RemoveTask(id uint64)
	RemoveAllTasks()
}
