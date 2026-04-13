package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Password  string         `gorm:"not null" json:"-"` // never return password in JSON
	FullName  string         `gorm:"not null" json:"full_name" validate:"required"`
	Role      string         `gorm:"default:customer" json:"role"` // customer or restaurant_admin
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	// DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Enables Soft Deletes
}
