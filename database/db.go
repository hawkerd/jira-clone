package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hawkerd/jira-clone/models"
)

var DB *gorm.DB // database to hold tasks, projects, etc

func Init() {
	// postgres string
	dsn := "host=localhost user=dhawk password=125861ford dbname=jira_clone port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	// automatically migrate structs
	DB.AutoMigrate(&models.Task{}, &models.Project{})
}
