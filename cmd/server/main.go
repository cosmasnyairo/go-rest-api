package main

import "fmt"

// Responsible for instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting our application")
	return nil
}
func main() {
	fmt.Println("Go REST APP")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
