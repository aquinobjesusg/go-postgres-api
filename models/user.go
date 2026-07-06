package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // Agrega ID, CreatedAt, UpdatedAt, DeletedAt automáticamente
	Name       string `gorm:"size:100;not null" json:"name"`
	Email      string `gorm:"uniqueIndex;size:100;not null" json:"email"`
}
