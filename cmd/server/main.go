package main

import (
	"context"
	"fmt"

	"github.com/cosmasnyairo/go-rest-api/internal/comment"
	"github.com/cosmasnyairo/go-rest-api/internal/db"
)

var (
	SuccessfullyConnectedToDB = "successfully to connect to database"
)

// Responsible for instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting our application")
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	if err := db.MigrateDB(); err != nil {
		return err
	}

	fmt.Println(SuccessfullyConnectedToDB)

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"1cd40daa-60da-66fa-61cd-66fabd61cd40",
	))

	return nil
}
func main() {
	fmt.Println("Go REST APP")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
