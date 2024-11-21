package models

import (
	"gorm.io/gorm"
)

type Election struct {
	gorm.Model
	Title     string     `json:"title"`
	Status    string     `json:"status"`
	WinnerID  uint       `json:"winner_id"`
	Aspirants []Aspirant `gorm:"foreignKey:ElectionID" json:"aspirants"`
}
