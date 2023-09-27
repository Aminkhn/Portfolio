package model

import "time"

// portfolio contains (skills & projects)
type Portfolio struct {
	Id       int    `db:"id" json:"id"`
	Content  string `db:"content" json:"content"`
	UserID   uint   `db:"userID" json:"userID"`
	Skills   []Technology
	PortProj []Project
}

// projects done till now
type Project struct {
	Id         int       `db:"id" json:"id"`
	Title      string    `db:"title" json:"title"`
	Content    string    `db:"content" json:"content"`
	Image      string    `db:"image" json:"image"`
	IsDone     bool      `db:"is_done" json:"is_done"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Technology []Technology
}

// technologies
type Technology struct {
	Id      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	Image   string `db:"image" json:"image"`
}
