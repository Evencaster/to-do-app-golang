package mysql_repo

import (
	"github.com/Evencaster/to-do-app-golang/entities"
	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model
	ID 		int
	Name 	string
}

func (t *TaskModel) toEntity() entities.Task {
	return entities.Task{
		Name: t.Name,
		ID: uint64(t.ID),
		Timestamp: t.CreatedAt.Unix(),
	}
}
