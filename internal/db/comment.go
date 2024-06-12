package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Dmitrylolo/go-rest-api/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func ConvertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var commentRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author 
		FROM comments 
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&commentRow.ID, &commentRow.Slug, &commentRow.Body, &commentRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}

	return ConvertCommentRowToComment(commentRow), nil
}
