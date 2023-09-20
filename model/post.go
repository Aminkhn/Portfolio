package model

import "time"

type Post struct {
	Id           int       `db:"id" json:"id"`
	Title        string    `db:"title" json:"title"`
	Content      string    `db:"content" json:"content"`
	IsPublished  bool      `db:"is_published" json:"is_published"`
	Created_at   time.Time `db:"created_at" json:"created_at"`
	Updated_at   time.Time `db:"updated_at" json:"updated_at"`
	Published_at time.Time `db:"published_at" json:"published_at"`
	UserID       uint      `db:"userID" json:"userID"`
}

// every post has many comment
type Comment struct {
	Id           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	Content      string    `db:"content" json:"content"`
	IsPublished  bool      `db:"is_published" json:"is_published"`
	Created_at   time.Time `db:"created_at" json:"created_at"`
	Published_at time.Time `db:"published_at" json:"published_at"`
	PostID       uint      `db:"postID" json:"postID"`
}
