package main

import "fmt"

// Run is going to be responsible
// for the instantiation and sturtup of the application
func Run() error {
	fmt.Println("Starting application")
	return nil
}

func main() {
	// TODO: Implement main function
	fmt.Println("Hello, World!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
