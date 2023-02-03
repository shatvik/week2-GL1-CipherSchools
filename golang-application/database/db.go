package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/shatvik/week2-GL1-CipherSchools.git/models"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "postgres"
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+dbName+" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}
