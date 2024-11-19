package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserID     uint `json:"user_id" gorm:"not null"`
	ElectionID uint `json:"election_id" gorm:"not null"`
	AspirantID uint `json:"aspirant_id" gorm:"not null"`
}
