package models

import "gorm.io/gorm"

type Vote struct {
    gorm.Model
    AspirantID uint   `json:"aspirant_id"`
    ElectionID uint   `json:"election_id"`
    Aspirant   Aspirant `gorm:"foreignKey:AspirantID"`
    Election   Election `gorm:"foreignKey:ElectionID"`
}
