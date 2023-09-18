package model

import "time"

// portfolio contains (skills & projects)
type Portfolio struct {
	Id          int    `db:"id"`
	Description string `db:"description"`
	UserID      uint
	Skills      []Technology
	PortProj    []Project
}

// projects done till now
type Project struct {
	Id         int       `db:"id"`
	Title      string    `db:"title"`
	Text       string    `db:"text"`
	Image      string    `db:"image"`
	IsDone     bool      `db:"is_done"`
	Created_at time.Time `db:"created_at"`
	Technology []Technology
}

// technologies
type Technology struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Image       string `db:"image"`
}
