package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Stettzy/blog_in_golang/db"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	fmt.Print(password)
	fmt.Print("\n")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Print(hashedPassword)

	return string(hashedPassword), err
}

func (u *User) CreateUser() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(createUser)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hash, err := HashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(u.Email, hash, u.Username)
	if err != nil {
		return 0, err
	}

	r, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func (u *User) UpdateUser() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(updateUser)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Email, u.Username, u.Password)
	if err != nil {
		return 0, err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(r), err
}

func (u *User) DeleteUser() (int, error) {
	db, err := db.Get()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare(deleteUser)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.ID)
	if err != nil {
		return 0, err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(r), nil
}

func GetById(id int) (*User, error) {
	db, err := db.Get()
	if err != nil {
		return nil, err
	}

	stmt := db.QueryRow(getUserById, id)

	var user User

	err = stmt.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no user found")
		}

		return nil, errors.New("error scanning database columns")
	}

	return &user, nil
}

func GetByEmail(email string) (*User, error) {
	db, err := db.Get()
	if err != nil {
		return nil, err
	}

	stmt := db.QueryRow(getUserByEmail, email)

	var user User

	err = stmt.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no user found")
		}

		return nil, errors.New("error scanning database columns")
	}

	return &user, nil
}

const createUser = `INSERT INTO users (email, password, username) VALUES (?, ?, ?)`
const updateUser = `UPDATE users (emai, password, username) VALUES (?, ?, ?) WHERE id = ?`
const deleteUser = `DELETE FROM users WHERE id = ?`
const getUserByEmail = `SELECT id, email, username, password FROM users where email = ?`
const getUserById = `SELECT id, email, username, password FROM users where id = ?`
