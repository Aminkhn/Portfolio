package logic

import (
	"github.com/aminkhn/portfolio/model"
	"github.com/jmoiron/sqlx"
)

type PortfolioQuery struct {
	*sqlx.DB
}

func (q *PortfolioQuery) GetAllPortolios() ([]model.Portfolio, error) {

	files := []model.Portfolio{}
	//main portfolio
	PortfolioQuery := `SELECT * FROM Portfolio;`
	err := q.Select(&files, PortfolioQuery)
	if err != nil {
		return files, err
	}

	var skill []model.Technology
	var projects []model.Project

	for i, v := range files {
		// projects
		ProjectQuery := `SELECT  Project.id as id, Project.title as title, Project.contect as content, Project.image as image, Project.is_done as is_done, Project.created_at as created_at
			from portfolio_projects
			LEFT JOIN Project ON Project.id = portfolio_projects.projectID WHERE portfolio_projects.portfolioID = ?`
		err = q.Select(&projects, ProjectQuery, v.Id)
		if err != nil {
			return files, err
		}
		// skills
		TechsQuery := `SELECT   Technology.id as id, Technology.title as title, Technology.content as content, Technology.image as image
			from portfolio_techs
			LEFT JOIN Technology ON Technology.id = portfolio_techs.technologyID WHERE portfolio_techs.portfolioID = ?`
		err = q.Select(&skill, TechsQuery, v.Id)
		if err != nil {
			return files, err
		}
		// assigning
		files[i].Skills = skill
		files[i].PortProj = projects
	}
	return files, nil
}

func (q *PortfolioQuery) GetPortfolio(id string) (model.Portfolio, error) {
	//main portfolio
	file := model.Portfolio{}
	PortfolioQuery := `SELECT * FROM Portfolio WHERE id = ?;`
	err := q.Get(&file, PortfolioQuery, id)
	if err != nil {
		return file, err
	}

	var skill []model.Technology
	var projects []model.Project

	// projects
	ProjectQuery := `SELECT  Project.id as id, Project.title as title, Project.contect as content, Project.image as image, Project.is_done as is_done, Project.created_at as created_at
			from portfolio_projects
			LEFT JOIN Project ON Project.id = portfolio_projects.projectID WHERE portfolio_projects.portfolioID = ?`
	err = q.Select(&projects, ProjectQuery, id)
	if err != nil {
		return file, err
	}
	// skills
	TechsQuery := `SELECT   Technology.id as id, Technology.title as title, Technology.content as content, Technology.image as image
			from portfolio_techs
			LEFT JOIN Technology ON Technology.id = portfolio_techs.technologyID WHERE portfolio_techs.portfolioID = ?`
	err = q.Select(&skill, TechsQuery, id)
	if err != nil {
		return file, err
	}
	// assigning
	file.Skills = skill
	file.PortProj = projects

	return file, nil
}

// create new user
func (q *PortfolioQuery) CreatePortfolio(portfolio model.Portfolio) error {

	// func => take user id
	// create new portfolio record
	portfolio.UserID = 2
	queryPortfolio := `INSERT INTO portfolio (
		content, userID)
		VALUES (
		?, ?
	  );`
	res, err := q.Exec(queryPortfolio,
		portfolio.Content,
		portfolio.UserID,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	// create interface relation records
	queryProject := `INSERT INTO portfolio_projects (
		portfolioID, projectID
	) VALUES (
		?, ?
	);`
	for _, v := range portfolio.PortProj {
		_, err := q.Exec(queryProject,
			id,
			v.Id,
		)
		if err != nil {
			return err
		}
	}
	// create interface relation records
	querySkill := `INSERT INTO portfolio_techs (
		portfolioID, technologyID
	) VALUES (
		?, ?
	);`
	for _, v := range portfolio.Skills {
		_, err := q.Exec(querySkill,
			id,
			v.Id,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

/***********************************************/
// update existing main portfolio
func (q *PortfolioQuery) UpdatePortfolio(portfolio model.Portfolio, id string) error {

	query := `UPDATE portfolio
		set 
		description = ?
		WHERE id = ?;`

	_, err := q.Exec(query,
		portfolio.Content,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

// update existing projects of portfolio
func (q *PortfolioQuery) UpdateProjectsOfPortfolio(portfolio model.Portfolio, id string) error {

	query := `UPDATE portfolio
		set 
		title = ?, content = ?, user = ?
		WHERE id = ?;`

	_, err := q.Exec(query,
		portfolio.Content,
		portfolio.UserID,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

// update existing skills of portfolio
func (q *PortfolioQuery) UpdateSkillsOfPortfolio(portfolio model.Portfolio, id string) error {

	query := `UPDATE portfolio
		set 
		title = ?, content = ?, user = ?
		WHERE id = ?;`

	_, err := q.Exec(query,
		portfolio.Content,
		portfolio.UserID,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

/*********************************************/
// delete portfolio by id
func (q *PortfolioQuery) DeletePortfolio(id string) error {

	queryPorj := `DELETE FROM portfolio_projects WHERE portfolioID = ? ;`
	queryTech := `DELETE FROM portfolio_techs WHERE portfolioID = ? ;`
	queryPort := `DELETE FROM portfolio WHERE id = ? ;`
	q.MustExec(queryPorj, id)
	q.MustExec(queryTech, id)
	q.MustExec(queryPort, id)
	err := q.MustBegin().Commit()
	if err != nil {
		q.MustBegin().Rollback()
		return err
	}
	return nil
}
