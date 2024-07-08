package comment

import "github.com/Stettzy/blog_in_golang/db"

func (c *Comment) CreateComment() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(createComment)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(c.Content)
	if err != nil {
		return 0, err
	}

	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (c *Comment) UpdateComment() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(updateComment)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(c.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (c *Comment) LkeComment() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(likeComment)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(c.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (c *Comment) RemoveComment() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(removeComment)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(c.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

const createComment = `INSERT INTO comments (content, likes) VALUES (?, 0)`
const likeComment = `UPDATE comments SET id = id + 1 WHERE id = ?`
const updateComment = `UPDATE comments (content, likes) WHERE id = ?`
const removeComment = `REMOVE FROM comments WHERE id = ?`
