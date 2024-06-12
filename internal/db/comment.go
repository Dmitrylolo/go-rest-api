package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Dmitrylolo/go-rest-api/internal/comment"
	uuid "github.com/satori/go.uuid"
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

func (d *Database) CreateComment(ctx context.Context, c comment.Comment) (comment.Comment, error) {
	c.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     c.ID,
		Slug:   sql.NullString{String: c.Slug, Valid: true},
		Body:   sql.NullString{String: c.Body, Valid: true},
		Author: sql.NullString{String: c.Author, Valid: true},
	}

	row, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments 
		(id, slug, body, author) 
		VALUES (:id, :slug, :body, :author)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error creating comment: %w", err)
	}
	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing comment: %w", err)
	}

	return c, nil
}

func (d *Database) DeleteComment(ctx context.Context, uuid string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id = $1`,
		uuid,
	)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, c comment.Comment) (comment.Comment, error) {
	commentRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: c.Slug, Valid: true},
		Body:   sql.NullString{String: c.Body, Valid: true},
		Author: sql.NullString{String: c.Author, Valid: true},
	}

	row, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments 
		SET slug = :slug, body = :body, author = :author 
		WHERE id = :id`,
		commentRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating comment: %w", err)
	}
	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing comment: %w", err)
	}
	return ConvertCommentRowToComment(commentRow), nil
}
