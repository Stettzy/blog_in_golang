package migrations

import (
	"fmt"

	"github.com/Stettzy/blog_in_golang/db"
)

func CreateUsers() error {
	db, err := db.Get()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50) NOT NULL,
        email VARCHAR(100) NOT NULL,
        password VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}
