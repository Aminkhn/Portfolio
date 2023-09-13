package model

import (
	"time"

	"gorm.io/gorm"
)

// portfolio contains (skills & projects)
type Portfolio struct {
	gorm.Model
	UserID      uint
	Technology  []Technology `json:"technology" gorm:"many2many:portfo_techs"`
	Project     []Project    `json:"project" gorm:"many2many:portfo_projects"`
	Description string       `json:"description"`
}

// projects done till now
type Project struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
	//Technology []Technology `json:"technology" gorm:"many2many:project_techs"`
	Image      string    `json:"image"`
	IsDone     bool      `json:"is_done" gorm:"default:false"`
	Created_at time.Time `json:"created_at"`
}

// technologies
type Technology struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
