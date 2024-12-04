package models

import "gorm.io/gorm"

// Initialize Users model
type Users struct {
	gorm.Model
	Username string `gorm:"unique;size:255;not null" json:"Username,omitempty"`
	Password string `gorm:"size:255;not null"  json:"Password,omitempty"`
	Role     string `gorm:"size:10;not null;default:'user'" json:"Role,omitempty"`
}
