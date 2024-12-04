package models

import "gorm.io/gorm"

// Initialize Users model
type Users struct {
	gorm.Model
	Username string `gorm:"unique;size:255;not null" json:"username,omitempty"`
	Password string `gorm:"size:255;not null"  json:"password,omitempty"`
	Role     string `gorm:"size:10;not null;default:'user'" json:"role,omitempty"`
}
