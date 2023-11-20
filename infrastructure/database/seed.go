package database

import (
	"errors"
	"final_project_3/entity"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var admin entity.User
	err := db.Where("id = ? AND full_name = ?", "1", "admin").First(&admin)
	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			admin = entity.User{FullName: "admin", Email: "admin@mail.com", Password: "admin123", Role: "admin"}
			db.Create(&admin)
		}
	}
}