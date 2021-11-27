package repository

import (
	"context"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/pkg/errors"
)

type taskTypeRepository struct {
	db *ent.Client
}

func NewTaskTypeRepository(db *ent.Client) domain.TaskTypeRepository {
	return &taskTypeRepository{db: db}
}

func (r *taskTypeRepository) Save(targetID int, taskType *ent.TaskType) (*ent.TaskType, error) {
	task, err := r.db.TaskType.Create().
		SetType(taskType.Type).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return task, nil
}
