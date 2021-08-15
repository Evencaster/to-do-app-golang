package repository

import (
	"sync"
	"time"
)

type Task struct {
	ID 			uint64		`json:"id,omitempty"`
	Name 		string		`json:"name"`
	Timestamp 	int64		`json:"timestamp"`
}

type Repo struct {
	m 			sync.RWMutex
	sequence 	uint64
	tasks 		[]Task
}

func New() *Repo {
	return &Repo{tasks: []Task{}}
}

func (r *Repo) GetAllTasks() []Task {
	r.m.RLock()
	defer r.m.RUnlock()

	return r.tasks
}

func (r *Repo) AddTask(name string) uint64 {
	r.m.Lock()
	defer r.m.Unlock()
	r.sequence++
	r.tasks = append(r.tasks, Task{
		Name: name,
		ID: r.sequence,
		Timestamp: time.Now().Unix(),
	})
	return r.sequence
}

func (r *Repo) RemoveTask(id uint64)  {
	r.m.Lock()
	defer r.m.Unlock()

	var newTasks []Task
	for i := 0; i < len(r.tasks); i++ {
		if r.tasks[i].ID != id {
			newTasks = append(newTasks, r.tasks[i])
		}
	}
	r.tasks = newTasks
}

func (r *Repo) RemoveAllTasks() {
	r.m.Lock()
	defer r.m.Unlock()

	r.tasks = []Task{}
}
