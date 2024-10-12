package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hawkerd/jira-clone/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// Get environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := "5432" // Set via env var or hardcode

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, dbname, port)

	var err error

	// Retry loop to connect to the database
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connected to the PostgreSQL database!")
			break
		}
		log.Printf("PostgreSQL connection failed: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	// Final error check
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL after several attempts: %v", err)
	}

	// Automigrate the models
	DB.AutoMigrate(&models.Task{}, &models.Project{})
}
