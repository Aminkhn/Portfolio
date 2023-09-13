package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Text         string    `json:"text"`
	IsPublished  bool      `json:"is_published"`
	Created_at   time.Time `json:"created_at"`
	Published_at time.Time `json:"published_at"`
	Comments     []Comment
	UserID       uint
}

// every post has many comment
type Comment struct {
	ID           int       `json:"id" gorm:"uniqueIndex; not null;"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Text         string    `json:"text"`
	IsPublished  bool      `json:"is_published" gorm:"default:false"`
	Published_at time.Time `json:"published_at"`
	Created_at   time.Time `json:"created_at"`
	PostID       uint
}
