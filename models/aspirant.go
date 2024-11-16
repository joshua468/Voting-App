package models

import (
	"gorm.io/gorm"
)

type Aspirant struct {
	gorm.Model
	Name     string `json:"name"`
	Party    string `json:"party"`
	Photo    string `json:"photo"`
	EctionID uint   `json:"election_id"`
	Votes    int    `gorm:default:0  json:"votes"`
}
