package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	VoterID  string `gorm:"unique" json:"voter_id"`
	Verified bool   `gorm:"default:false"  json:"verified"`
	Role     string `gorm:"default:user" json:"role"`
}
