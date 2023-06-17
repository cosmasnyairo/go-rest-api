package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cosmasnyairo/go-rest-api/internal/comment"
)

var (
	ErrorFetchingComment = "error fetching comment by uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Author sql.NullString
	Body   sql.NullString
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
		return comment.Comment{}, fmt.Errorf("%s: %w", ErrorFetchingComment, err)
	}

	return convertCommentRowtToComment(cmtRow), nil
}
