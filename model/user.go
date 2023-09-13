package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string    `json:"username" gorm:"varchar(50);uniqueIndex; not null"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Family      string    `json:"family"`
	Password    string    `json:"password"`
	Role        string    `gorm:"not null" json:"role"`
	Image       string    `json:"image"`
	PhoneNumber int       `json:"phone_number" gorm:"uniqueIndex"`
	Created_at  time.Time `json:"created_at"`
}
