package database

import (
	"fmt"
	"github.com/galantixa/go-gin-gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func SetupDb() {
	dns := "host=localhost user=postgres password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect Databases!", err)
	}

	DB.AutoMigrate(&models.Book{})
	fmt.Println("Databases migrated!")
}
