package post

import "github.com/Stettzy/blog_in_golang/db"

func (p *Post) CreatePost() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(createPost)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Title, p.Content)
	if err != nil {
		return 0, err
	}

	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (p *Post) UpdatePost() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(updatePost)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Title, p.Content, p.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (p *Post) RemovePost() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(removePost)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

const createPost = `INSERT INTO posts (title, content) VALUES (?, ?)`
const updatePost = `UPDATE posts (title, content) VALUES (?, ?) WHERE id = ?`
const removePost = `REMOVE FROM posts WHERE id = ?`
