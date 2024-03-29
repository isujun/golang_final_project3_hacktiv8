package task_repository

import (
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
)

type Repository interface {
	CreateTask(taskPayload *entity.Task) (*entity.Task, errs.Error)
	GetTasks() (*[]dto.GetTasksResponse, errs.Error)
	UpdateTask(taskPayload *entity.Task) (*entity.Task, errs.Error)
	UpdateStatus(taskPayload *entity.Task) (*entity.Task, errs.Error)
	UpdateTaskCategory(taskPayload *entity.Task) (*entity.Task, errs.Error)
	DeleteTask(id int) errs.Error
}
