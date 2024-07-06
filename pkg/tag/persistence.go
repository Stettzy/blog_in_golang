package tag

import "github.com/Stettzy/blog_in_golang/db"

func (t *Tag) Create() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(createTag)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (t *Tag) Update() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(updateTag)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (t *Tag) Remove() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(deleteTag)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.ID)

	if err != nil {
		return 0, err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

const createTag = `INSERT INTO tags (title) VALUES (?)`
const updateTag = `UPDATE tags (title) VALUES (?) where id = ?`
const deleteTag = `DELETE FROM tags where id = ?`
