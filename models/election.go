package models

import "gorm.io/gorm"

type Election struct {
    gorm.Model
    Title      string `json:"title"`
    StartDate  string `json:"start_date"`
    EndDate    string `json:"end_date"`
    Aspirants  []Aspirant `json:"aspirants"`
    Votes      []Vote `json:"votes"`
}
