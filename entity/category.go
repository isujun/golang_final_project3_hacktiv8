package entity

import "time"

type Category struct {
	ID        uint `gorm:"primaryKey;not null"`
	Type      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
