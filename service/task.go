package service

import (
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/pkg/helpers"
	task_repository "final_project_3/repository/task_repository"
)

type taskService struct {
	TaskRepo task_repository.Repository
}

type TaskService interface {
	CreateTask(taskPayload *dto.NewTaskRequest) (*dto.NewTaskResponse, errs.Error)
	GetTasks() (*[]dto.GetTasksResponse, errs.Error)
}

func NewTaskService(taskRepo task_repository.Repository) TaskService {
	return &taskService{TaskRepo: taskRepo}
}

func (ts *taskService) CreateTask(taskPayload *dto.NewTaskRequest) (*dto.NewTaskResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(taskPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	task := entity.Task{
		Title: taskPayload.Title,
		Description: taskPayload.Description,
		CategoryID: taskPayload.CategoryID,
		UserID: taskPayload.UserID,
	}
	createdTask, err := ts.TaskRepo.CreateTask(&task)
	if err != nil {
		return nil, err
	}
	response := dto.NewTaskResponse{
		ID: createdTask.ID,
		Title: createdTask.Title,
		Description: createdTask.Description,
		UserID: createdTask.UserID,
		CategoryID: createdTask.CategoryID,
		CreatedAt: createdTask.CreatedAt,
	}
	return &response, nil
}

func (ts *taskService) GetTasks() (*[]dto.GetTasksResponse, errs.Error) {
	response, err := ts.TaskRepo.GetTasks()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (ts *taskService) UpdateTask(taskPayload *dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(taskPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	var task = entity.Task{
		Title: taskPayload.Title,
		Description: taskPayload.Description,
	}
	updateTask, err := ts.TaskRepo.UpdateTask(&task)
	if err != nil {
		return nil, err
	}
	response := dto.UpdateTaskResponse{
		ID: updateTask.ID,
		Title: updateTask.Title,
		Description: updateTask.Description,
		Status: updateTask.Status,
		UserID: updateTask.UserID,
		CategoryID: updateTask.CategoryID,
		UpdatedAt: updateTask.UpdatedAt,
	}
	return &response, nil
}
