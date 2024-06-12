package main

import (
	"fmt"

	"github.com/Dmitrylolo/go-rest-api/internal/comment"
	"github.com/Dmitrylolo/go-rest-api/internal/db"
	transportHttp "github.com/Dmitrylolo/go-rest-api/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(commentService)
	if err := httpHandler.Serve(); err != nil {
		fmt.Println("Failed to start http server")
		return err
	}

	return nil
}

func main() {
	// TODO: Implement main function
	fmt.Println("Hello, World!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
