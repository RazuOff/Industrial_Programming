package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get

func initDB() {
	dsn := "host=45.10.41.15 user=Makism password=root dbname=MaksimShop port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	db.AutoMigrate(&testProducts)
	db.AutoMigrate(&testUsers)
}
