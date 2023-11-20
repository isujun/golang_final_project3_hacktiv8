package category_repo

import (
	"errors"
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/repository/category_repository"
	"strconv"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) category_repository.Repository {
	return &categoryRepo{db: db}
}

func (cr *categoryRepo) CreateCategory(categoryPayload *entity.Category) (*entity.Category, errs.Error) {
	var Category = *categoryPayload
	err := cr.db.Create(&Category).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &Category, nil
}

func (cr *categoryRepo) GetCategoriesWithTasks() (*[]dto.GetCategoriesResponse, errs.Error) {
	var categories []entity.Category
	err := cr.db.Find(&categories).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	var res []dto.GetCategoriesResponse
	for _, category := range categories {
		var tasks []entity.Task
		var taskDto []dto.GetCategoriesTaskResponse
		err = cr.db.Where("category_id = ?", category.ID).Find(&tasks).Error
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		if len(tasks) > 0 {
			for _, task := range tasks {
				taskDto = append(taskDto, dto.GetCategoriesTaskResponse{
					ID: task.ID,
					Title: task.Title,
					Description: task.Description,
					UserID: task.UserID,
					CategoryID: task.CategoryID,
					CreatedAt: task.CreatedAt,
					UpdatedAt: task.UpdatedAt,
				})
			}
		}
		res = append(res, dto.GetCategoriesResponse{
			ID:        category.ID,
			Type:      category.Type,
			UpdatedAt: category.UpdatedAt,
			CreatedAt: category.CreatedAt,
			Task:      taskDto,
		})
	}
	return &res, nil
}

func (cr *categoryRepo) UpdateCategory(categoryPayload *entity.Category) (*entity.Category, errs.Error) {
	var Category = *categoryPayload
	err := cr.db.Model(&Category).Update("type", Category.Type).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &Category, nil
}

func (cr *categoryRepo) DeleteCategory(id int) errs.Error {
	err := cr.db.Where("id = ?", id).Delete(&entity.Category{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "todo with id: " + strconv.Itoa(id) + " not found"
			return errs.NewNotFoundError(msg)
		}
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}