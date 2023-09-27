package logic

import (
	"time"

	"github.com/aminkhn/portfolio/model"
	"github.com/jmoiron/sqlx"
)

type PostQuery struct {
	*sqlx.DB
}

func (q *PostQuery) GetAllPosts() ([]model.Post, error) {

	files := []model.Post{}
	query := `select * from post`

	err := q.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, nil
}

func (q *PostQuery) GetOnePost(id string) (model.Post, error) {

	file := model.Post{}
	query := `select * from post where id = ?`

	err := q.Get(&file, query, id)
	if err != nil {
		return file, err
	}
	return file, nil
}

// create new user
func (q *PostQuery) CreatePost(post model.Post) error {

	query := `INSERT INTO post (
		title, content, is_published, created_at, updated_at, published_at, userID)
		VALUES (
		?, ?, ?, ?, ?, ?, ?
	  );`

	post.Created_at = time.Now()
	post.Updated_at = time.Now()
	post.Published_at = time.Now()
	post.IsPublished = false

	// func => take user id
	post.UserID = 2

	_, err := q.Exec(query,
		post.Title,
		post.Content,
		post.IsPublished,
		post.Created_at,
		post.Updated_at,
		post.Published_at,
		post.UserID,
	)

	if err != nil {
		return err
	}
	return nil
}

// update existing user
func (q *PostQuery) UpdatePost(post model.Post, id string) error {

	query := `UPDATE post
		set 
		title = ?, content = ?, updated_at = ?
		WHERE id = ?;`

	if post.IsPublished {
		post.Published_at = time.Now()
	}
	post.Updated_at = time.Now()
	_, err := q.Exec(query,
		post.Title,
		post.Content,
		post.Updated_at,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

// delete user by id
func (q *PostQuery) DeletePost(id string) error {

	query := `DELETE FROM post WHERE id = ?;`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

/***************************************************/
func (q *PostQuery) GetPublishedPosts() ([]model.Post, error) {

	files := []model.Post{}
	query := `select * from post where is_published = true`

	err := q.Select(&files, query)
	if err != nil {
		return files, err
	}
	return files, nil
}
func (q *PostQuery) PublishPost(post model.Post, id string) error {

	query := `update post set is_published = ?, published_at = ? where id = ?`

	published_at := time.Now()
	_, err := q.Exec(query, post.IsPublished, published_at, id)
	if err != nil {
		return err
	}
	return nil
}
