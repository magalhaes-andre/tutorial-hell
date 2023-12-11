package database

import (
	"log"

	"github.com/magalhaes-andre/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error while attempting to connect with database: ", err)
	}
	DB.AutoMigrate(&models.User{})
}
