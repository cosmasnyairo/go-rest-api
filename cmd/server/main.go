package main

import (
	"context"
	"fmt"

	"github.com/cosmasnyairo/go-rest-api/internal/comment"
	"github.com/cosmasnyairo/go-rest-api/internal/db"
)

var (
	SuccessfullyConnectedToDB = "successfully connected to database"
)

// Responsible for instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting application")
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	if err := db.MigrateDB(); err != nil {
		return err
	}

	fmt.Println(SuccessfullyConnectedToDB)

	cmtService := comment.NewService(db)

	newcomment := comment.Comment{
		ID:     "1cd40daa-60da-66fa-61cd-66fabd61cd40",
		Slug:   "manual-test",
		Body:   "Cosmas",
		Author: "Hola! Soy Dora!",
	}

	updatedcomment := comment.Comment{
		Slug:   "not-a-manual-test",
		Body:   "Cosmas",
		Author: "Hola! Soy Dora!",
	}

	createdcomment, _ := cmtService.CreateComment(context.Background(), newcomment)
	fmt.Println(cmtService.GetComment(context.Background(), createdcomment.ID))

	cmtService.UpdateComment(context.Background(), createdcomment.ID, updatedcomment)
	fmt.Println(cmtService.GetComment(context.Background(), createdcomment.ID))

	cmtService.DeleteComment(context.Background(), createdcomment.ID)
	fmt.Println(cmtService.GetComment(context.Background(), createdcomment.ID))

	return nil
}
func main() {
	fmt.Println("Go REST APP")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
