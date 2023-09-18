package model

type Management struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Meta        string `db:"meta"`
	Description string `db:"desc"`
	Author      string `db:"author"`
	Phone       string `db:"phone"`
	Email       string `db:"email"`
}
