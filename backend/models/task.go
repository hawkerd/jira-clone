package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model          // id, created/updated/deleted at
	Title       string  `gorm:"not null"` // title (required)
	Description string  // description
	Status      string  `gorm:"not null"` // status (required)
	ProjectID   uint    // associated project id (required)
	Project     Project `gorm:"foreignKey:ProjectID"` // defines task-project relationship in database
}
