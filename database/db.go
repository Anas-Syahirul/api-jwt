package database

import (
	"api-jwt/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	dbPort   = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "db2"
	db       *gorm.DB
	err      error
)

func StartDB() {
	log.Println("connecting to database...")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	log.Println("successfully connect to database")

	db.AutoMigrate(models.User{}, models.Product{})

	log.Println("successfully migrate database")
}

func GetDB() *gorm.DB {
	return db
}
