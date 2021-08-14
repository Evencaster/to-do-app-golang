package repository

import (
	"sync"
	"time"
)

type Task struct {
	ID 			int64		`json:"id,omitempty"`
	Name 		string		`json:"name"`
	Timestamp 	int64		`json:"timestamp"`
}

type Repo struct {
	sequence 	int64
	m 			sync.Mutex
	tasks 		[]Task
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetAllTasks() []Task {
	r.m.Lock()
	defer r.m.Unlock()

	return r.tasks
}

func (r *Repo) AddTask(name string) int64 {
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

func (r *Repo) RemoveTask(id int64)  {
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
