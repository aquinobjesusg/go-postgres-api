package models

import "gorm.io/gorm"

// User representa la tabla en PostgreSQL
type User1 struct {
	gorm.Model        // Agrega ID, CreatedAt, UpdatedAt, DeletedAt automáticamente
	Name       string `json:"name" gorm:"not null"`
	Email      string `json:"email" gorm:"unique;not null"`
}
