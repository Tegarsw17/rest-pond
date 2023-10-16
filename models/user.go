package models

import "time"

type Users struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password" gorm:"unique"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
