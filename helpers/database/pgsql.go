package pgsql

import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB(envConfig models.Database) {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", envConfig.Host, envConfig.User, envConfig.Password, envConfig.Port, envConfig.Name)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{}, models.Brand{}, models.Car{})
}

func GetDB() *gorm.DB {
	return db
}
