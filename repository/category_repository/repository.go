package category_repository

import (
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
)

type Repository interface {
	CreateCategory(categoryPayload *entity.Category) (*entity.Category, errs.Error)
	GetCategoriesWithTasks() (*[]dto.GetCategoriesResponse, errs.Error)
	UpdateCategory(categoryPayload *entity.Category) (*entity.Category, errs.Error)
	DeleteCategory(id int) errs.Error
}
