package mem_repo

import (
	"github.com/Evencaster/to-do-app-golang/repository"
	"sync"
	"time"
)

type MemRepo struct {
	m        sync.RWMutex
	sequence uint64
	tasks    []repo.Task
}

func NewMemRepo() *MemRepo {
	return &MemRepo{tasks: []repo.Task{}}
}

func (r *MemRepo) GetAllTasks() []repo.Task {
	r.m.RLock()
	defer r.m.RUnlock()

	return r.tasks
}

func (r *MemRepo) AddTask(name string) uint64 {
	r.m.Lock()
	defer r.m.Unlock()

	r.sequence++
	id := r.sequence
	r.tasks = append(r.tasks, repo.Task{
		Name:      name,
		ID:        id,
		Timestamp: time.Now().Unix(),
	})
	return r.sequence
}

func (r *MemRepo) RemoveTask(id uint64) {
	r.m.Lock()
	defer r.m.Unlock()

	var newTasks []repo.Task
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

	r.tasks = []repo.Task{}
}
