package models

import (
	"gorm.io/gorm"
)

type Aspirant struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null"`
	ElectionID uint   `json:"election_id" gorm:"not null"`
	Votes      int    `json:"votes" gorm:"default:0"`
}
