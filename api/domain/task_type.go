package domain

import "github.com/Wanted-Linx/linx-backend/api/ent"

type TaskTypeRequest struct {
	TaskType []string `json:"task_type"`
}

type TaskTypeRepository interface {
	Save(targetID int, taskType *ent.TaskType) (*ent.TaskType, error)
}
