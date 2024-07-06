package main

import (
	"fmt"
	"log"

	"github.com/Stettzy/blog_in_golang/db"
	"github.com/Stettzy/blog_in_golang/db/migrations"
	"github.com/Stettzy/blog_in_golang/handlers"
	"github.com/labstack/echo"
)

func RunMigrations() error {
	migrations := []struct {
		name string
		fn   func() error
	}{
		{"users", migrations.CreateUsers},
		{"posts", migrations.CreatePosts},
		{"tags", migrations.CreateTags},
		{"comments", migrations.CreateComments},
	}

	for _, migration := range migrations {
		if err := migration.fn(); err != nil {
			return fmt.Errorf("failed migration for %s: %w", migration.name, err)
		}
	}

	return nil
}

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = RunMigrations()

	if err != nil {
		log.Fatal("failed to execute: %w", err)
	}

	api := echo.New()

	// User routes
	api.POST("/login", handlers.LoginUser)
	api.POST("/register", handlers.RegisterUser)
	api.DELETE("/remove", handlers.RemoveUser)
	api.PUT("/update", handlers.UpdateUser)
	// Post routes
	// Tags routes
	// Comments routes

	api.Start(":8080")
}
