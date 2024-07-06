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
	api.PUT("/update", handlers.UpdateUser)
	api.POST("/register", handlers.RegisterUser)
	api.DELETE("/remove", handlers.RemoveUser)
	// Post routes
	api.POST("/post/create", handlers.CreatePost)
	api.PUT("/post/update", handlers.UpdatePost)
	api.DELETE("/post/remove", handlers.RemovePost)
	// Tags routes
	api.POST("/tag/create", handlers.CreateTag)
	api.PUT("/tag/update", handlers.UpdateTag)
	api.DELETE("/tag/remove", handlers.RemoveTag)
	// Comments routes
	api.POST("/comment/create", handlers.CreateComment)
	api.PUT("/comment/update", handlers.UpdateComment)
	api.DELETE("/comment/remove", handlers.RemoveComment)

	api.Start(":8080")
}
