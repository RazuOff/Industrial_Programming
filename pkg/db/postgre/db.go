package postgre

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host, exists := os.LookupEnv("HOST")
	if !exists {
		log.Fatal("Failed to connect to database:")
	}
	user, exists := os.LookupEnv("USER")
	if !exists {
		log.Fatal("Failed to connect to database:")
	}
	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		log.Fatal("Failed to connect to database:")
	}
	dbname, exists := os.LookupEnv("DBNAME")
	if !exists {
		log.Fatal("Failed to connect to database:")
	}
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("Failed to connect to database:")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	DB.AutoMigrate(&testProducts)
	DB.AutoMigrate(&testUsers)

	insertTestData()
}

func insertTestData() {
	DB.Create(testProducts)
	DB.Create(testUsers)
}
