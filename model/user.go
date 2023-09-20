package model

import "time"

type User struct {
	Id          int       `db:"id" json:"id"`
	Username    string    `db:"username" json:"username"`
	Email       string    `db:"email" json:"email"`
	Name        string    `db:"name" json:"name"`
	Family      string    `db:"family" json:"family"`
	Password    string    `db:"password" json:"password"`
	Role        string    `db:"role" json:"role"`
	Image       string    `db:"image" json:"image"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Created_at  time.Time `db:"created_at" json:"created_at"`
	Updated_at  time.Time `db:"updated_at" json:"updated_at"`
}
