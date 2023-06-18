package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cosmasnyairo/go-rest-api/internal/comment"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrorFetchingComment        = "error fetching comment by uuid"
	ErrorCreatingComment        = "failed to insert comment into database"
	ErrorDeletingComment        = "failed to delete comment from database"
	ErrorClosingRows            = "failed to close rows"
	ErrorUpdatingComment        = "failed to update comment"
	SuccessfullyInsertedComment = "comment inserted successfully"
	SuccessfullyUpdatedComment  = "comment updated successfully "
	SuccessfullyDeletedComment  = "comment deleted successfully "
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowtToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author
		FROM comments 
		WHERE id=$1`,
		uuid,
	)
	if err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author); err != nil {
		return comment.EmptyComment, fmt.Errorf("%s: %w", ErrorFetchingComment, err)
	}

	return convertCommentRowtToComment(cmtRow), nil
}

func (d *Database) CreateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	createdRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT into comments
		(id,slug, body, author)
		VALUES
		(:id, :slug, :body, :author)`,
		createdRow,
	)
	if err != nil {
		return comment.EmptyComment, fmt.Errorf("%s: %w", ErrorCreatingComment, err)
	}
	if err := rows.Close(); err != nil {
		return comment.EmptyComment, fmt.Errorf("%s: %w", ErrorClosingRows, err)
	}
	fmt.Printf("%s with id: %v\n", SuccessfullyInsertedComment, createdRow.ID)
	return convertCommentRowtToComment(createdRow), nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id=$1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrorDeletingComment, err)
	}
	fmt.Printf("%s with id: %v\n", SuccessfullyDeletedComment, id)

	return err
}

func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	createdRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug, 
		body = :body, 
		author = :author
		WHERE id = :id`,
		createdRow,
	)

	if err != nil {
		return comment.EmptyComment, fmt.Errorf("%s: %w", ErrorUpdatingComment, err)
	}
	if err := rows.Close(); err != nil {
		return comment.EmptyComment, fmt.Errorf("%s: %w", ErrorClosingRows, err)
	}
	fmt.Printf("%s with id: %v\n", SuccessfullyUpdatedComment, createdRow.ID)
	return convertCommentRowtToComment(createdRow), nil

}
