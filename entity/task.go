package entity

import "time"

type Task struct {
	ID          uint `gorm:"primaryKey;not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      bool `gorm:"not null"`
	UserID      uint
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User		User
	Category	Category
}

