package dto

import (
	"time"
)

type NewCategoryRequest struct {
	Type string `json:"type" validate:"required"`
}

type NewCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCategoriesResponse struct {
	ID        uint                        `json:"id"`
	Type      string                      `json:"type"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
	Task      []GetCategoriesTaskResponse `json:"Tasks"`
}

type GetCategoriesTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateCategoryRequest struct {
	ID   uint   `json:"id"`
	Type string `json:"type" validate:"required"`
}

type UpdateCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}
