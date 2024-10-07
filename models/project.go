package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model        // id, created/updated/deleted at
	Name       string `gorm:"not null"` // name (required)
}
