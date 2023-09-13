package model

type WebSite struct {
	Title       string `json:"title"`
	Meta        string `json:"meta"`
	Description string `json:"desc"`
	Author      string `json:"author"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
