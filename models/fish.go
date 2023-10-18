package models

import "time"

type Fish struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	Type            string    `json:"type"`
	Colour          string    `json:"colour"`
	Size            string    `json:"size"`
	Maintenance     string    `json:"maintenance"`
	DateMaintenance time.Time `json:"date_maintenance"`
	IdPond          string    `json:"id_pond"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CreatedFish struct {
	Type            string `json:"type"`
	Colour          string `json:"colour"`
	Size            string `json:"size"`
	Maintenance     string `json:"maintenance"`
	DateMaintenance string `json:"date_maintenance"`
	IdPond          string `json:"id_pond"`
}

type ShowFish struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Colour          string    `json:"colour"`
	Size            string    `json:"size"`
	Maintenance     string    `json:"maintenance"`
	DateMaintenance time.Time `json:"date_maintenance"`
	IdPond          string    `json:"id_pond"`
}
