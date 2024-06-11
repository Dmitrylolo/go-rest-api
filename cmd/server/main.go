package main

import (
	"context"
	"fmt"

	"github.com/Dmitrylolo/go-rest-api/internal/comment/db"
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

	if err := db.Ping(context.Background()); err != nil {
		fmt.Println("Failed to ping database")
		return err
	}
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
