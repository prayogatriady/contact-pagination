package model

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID        int            `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ContactResponse struct {
	ID        int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
