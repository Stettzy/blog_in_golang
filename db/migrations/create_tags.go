package migrations

import (
	"fmt"

	"github.com/Stettzy/blog_in_golang/db"
)

func CreateTags() error {
	db, err := db.Get()
	if err != nil {
		return fmt.Errorf("failed to get database %w", err)
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY,
		title TEXT	
	)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute statement %w", err)
	}

	return nil
}
