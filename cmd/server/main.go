package main

import (
	"context"
	"fmt"

	"github.com/Dmitrylolo/go-rest-api/internal/comment"
	"github.com/Dmitrylolo/go-rest-api/internal/db"
)

// Run is going to be responsible
// for the instantiation and sturtup of the application
func Run() error {
	fmt.Println("Starting application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.Migrate(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	commentService := comment.NewService(db)
	fmt.Println(commentService.GetComment(
		context.Background(),
		"9ae0a0d0-bd0c-4d1e-a3a4-d4e5f6a7b8c9",
	))

	fmt.Println("Connected to database")
	defer db.Client.Close()

	return nil
}

func main() {
	// TODO: Implement main function
	fmt.Println("Hello, World!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
