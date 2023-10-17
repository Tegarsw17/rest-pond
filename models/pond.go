package models

import (
	"time"
)

type Pond struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" validate:"required"`
	Dimension       string    `json:"dimension" validate:"required"`
	Condition       string    `json:"condition" binding:"required"`
	Maintenance     string    `json:"maintenance"`
	DateMaintenance time.Time `json:"dateMaintenance"`
	DateFeeding     time.Time `json:"dateFeeding"`
	TotalFish       int       `json:"totalFish"`
	IdUser          string    `json:"id_user"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ShowPond struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Dimension       string    `json:"dimension"`
	Condition       string    `json:"condition"`
	Maintenance     string    `json:"maintenance"`
	DateMaintenance time.Time `json:"dateMaintenance"`
	DateFeeding     time.Time `json:"dateFeeding"`
	TotalFish       int       `json:"totalFish"`
}
type CreatePond struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Dimension       string `json:"dimension"`
	Condition       string `json:"condition"`
	Maintenance     string `json:"maintenance"`
	DateMaintenance string `json:"dateMaintenance"`
	DateFeeding     string `json:"dateFeeding"`
	TotalFish       int    `json:"totalFish"`
}
