package service

import (
	"github.com/aminkhn/portfolio/model"
	"github.com/jmoiron/sqlx"
)

type PostQuery struct {
	*sqlx.DB
}

func (q *PostQuery) GetUsers() ([]model.Post, error) {

	files := []model.Post{}
	query := `select * from post`

	err := q.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, err
}

func (q *PostQuery) GetOnePost(id string) (model.Post, error) {

	file := model.Post{}
	query := `select * from post where id=$1`

	err := q.Get(&file, query, id)
	if err != nil {
		return file, err
	}
	return file, err
}
