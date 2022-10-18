package database

import (
	"assignment-two/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	user = "postgres"
	password = "dicky123"
	port = 5432
	dbname = "rest-api-go"
	db *gorm.DB
	err error
)

func StartDB()  {
	config := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)

	db ,err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(models.Orders{}, models.Items{})
}

func GetDb() *gorm.DB  {
	return db
}