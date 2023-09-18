package model

import "time"

type User struct {
	Id          int       `db:"id"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	Name        string    `db:"name"`
	Family      string    `db:"family"`
	Password    string    `db:"password"`
	Role        string    `db:"role"`
	Image       string    `db:"image"`
	PhoneNumber int       `db:"phone_number"`
	Created_at  time.Time `db:"created_at"`
}
