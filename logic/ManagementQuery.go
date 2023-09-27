package logic

import (
	"github.com/aminkhn/portfolio/model"
	"github.com/jmoiron/sqlx"
)

type ManagementQuery struct {
	*sqlx.DB
}

func (q *ManagementQuery) GetManagements() ([]model.Management, error) {

	files := []model.Management{}
	query := `select * from management`

	err := q.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, nil
}

func (q *ManagementQuery) UpdateManagement() (model.Management, error) {

	file := model.Management{}
	query := `select * from management where id = 1`

	err := q.Get(&file, query)
	if err != nil {
		return file, err
	}
	return file, nil
}
