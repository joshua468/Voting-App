package models

import "gorm.io/gorm"

type Election struct {
	gorm.Model
	Title     string     `json:"title" gorm:"not null"`
	StartDate string     `json:"start_date" gorm:"not null"`
	EndDate   string     `json:"end_date" gorm:"not null"`
	Aspirants []Aspirant `json:"aspirants" gorm:"foreignKey:ElectionID"`
}
