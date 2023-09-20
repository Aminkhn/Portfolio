package logic

import (
	"github.com/aminkhn/portfolio/model"
	"github.com/jmoiron/sqlx"
)

type PortfolioQuery struct {
	*sqlx.DB
}

func (q *PortfolioQuery) GetPortfolios() ([]model.Portfolio, error) {

	files := []model.Portfolio{}
	query := `select * from portfolio`

	err := q.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, err
}

func (q *PortfolioQuery) GetPortfolio(id string) (model.Portfolio, error) {

	file := model.Portfolio{}
	query := `select * from portfolio where id=$1`

	err := q.Get(&file, query, id)
	if err != nil {
		return file, err
	}
	return file, err
}
