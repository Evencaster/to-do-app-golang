package repository

import (
	"sync"
	"time"
)

type Task struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}

type MemRepo struct {
	m        sync.RWMutex
	sequence uint64
	tasks    []Task
}

func NewMemRepo() *MemRepo {
	return &MemRepo{tasks: []Task{}}
}

func (r *MemRepo) GetAllTasks() []Task {
	r.m.RLock()
	defer r.m.RUnlock()

	return r.tasks
}

func (r *MemRepo) AddTask(name string) uint64 {
	r.m.Lock()
	defer r.m.Unlock()

	r.sequence++
	id := r.sequence
	r.tasks = append(r.tasks, Task{
		Name:      name,
		ID:        id,
		Timestamp: time.Now().Unix(),
	})
	return r.sequence
}

func (r *MemRepo) RemoveTask(id uint64) {
	r.m.Lock()
	defer r.m.Unlock()

	var newTasks []Task
	for _, t := range r.tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		}
	}
	r.tasks = newTasks
}

func (r *MemRepo) RemoveAllTasks() {
	r.m.Lock()
	defer r.m.Unlock()

	r.tasks = []Task{}
}
