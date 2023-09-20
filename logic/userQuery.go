package logic

import (
	"time"

	"github.com/aminkhn/portfolio/model"
	"github.com/jmoiron/sqlx"
)

type UserQuery struct {
	*sqlx.DB
}

// list all users
func (q *UserQuery) GetAllUsers() ([]model.User, error) {

	files := []model.User{}
	query := `select * from user;`

	err := q.Select(&files, query)
	if err != nil {
		return []model.User{}, err
	}
	return files, nil
}

// get user by id
func (q *UserQuery) GetOneUser(id string) (model.User, error) {

	file := model.User{}
	query := `select * from user where id=?;`

	err := q.Get(&file, query, id)
	if err != nil {
		return model.User{}, err
	}
	return file, nil
}

// create new user
func (q *UserQuery) CreateUser(user model.User) error {

	query := `INSERT INTO user (
		username, email, name, family, password, role, image, phone_number, created_at, updated_at
	  ) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	  );`

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	HashPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = HashPassword
	_, err = q.Exec(query,
		user.Username,
		user.Email,
		user.Name,
		user.Family,
		user.Password,
		user.Role,
		user.Image,
		user.PhoneNumber,
		user.Created_at,
		user.Updated_at)

	if err != nil {
		return err
	}
	return nil
}

// update existing user
func (q *UserQuery) UpdateUser(user model.User, id string) error {

	query := `UPDATE user
		set 
		username = ?, email = ?, name = ?, family = ?, password = ?, role = ?,
		image = ?, phone_number = ?
		WHERE id = ?;`

	_, err := q.Exec(query,
		user.Username,
		user.Email,
		user.Name,
		user.Family,
		user.Password,
		user.Role,
		user.Image,
		user.PhoneNumber,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

// delete user by id
func (q *UserQuery) DeleteUser(id string) error {

	query := `DELETE FROM user WHERE id = ?;`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
