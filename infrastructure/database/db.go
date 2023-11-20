package database

import (
	"final_project_3/entity"
	"final_project_3/infrastructure/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func handleDatabaseConnection() (*gorm.DB, error) {
	appConfig := config.GetAppConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error occured while trying to connect to database", err)
	}

	db.Debug().AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Task{})
	return db, nil
}

func GetDatabaseInstance() *gorm.DB {
	db, err := handleDatabaseConnection()
	if err != nil {
		log.Panic(err)
	}
	return db
}
