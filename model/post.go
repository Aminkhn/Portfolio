package model

import "time"

type Post struct {
	Id           int       `db:"id"`
	Title        string    `db:"title"`
	Text         string    `db:"text"`
	IsPublished  bool      `db:"is_published"`
	Created_at   time.Time `db:"created_at"`
	Published_at time.Time `db:"published_at"`
	UserID       uint
}

// every post has many comment
type Comment struct {
	Id           int       `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	Text         string    `db:"text"`
	IsPublished  bool      `db:"is_published"`
	Published_at time.Time `db:"published_at"`
	Created_at   time.Time `db:"created_at"`
	PostID       uint
}
