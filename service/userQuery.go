package service

import (
	"github.com/aminkhn/portfolio/model"
	"github.com/jmoiron/sqlx"
)

type UserQuery struct {
	*sqlx.DB
}

func (q *UserQuery) GetAllUsers() ([]model.User, error) {

	files := []model.User{}
	query := `select * from user`

	err := q.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, err
}

func (q *UserQuery) GetOneUser(id string) (model.User, error) {

	file := model.User{}
	query := `select * from user where id=$1`

	err := q.Get(&file, query, id)
	if err != nil {
		return file, err
	}
	return file, err
}
