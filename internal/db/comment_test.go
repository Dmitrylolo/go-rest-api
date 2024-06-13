//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/Dmitrylolo/go-rest-api/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "new-author",
			Body:   "new-body",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "new-author",
			Body:   "new-body",
		})
		assert.NoError(t, err)

		newCmt, err := db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug:   "new-slug-2",
			Author: "new-author-2",
			Body:   "new-body-2",
		})
		assert.NoError(t, err)

		assert.Equal(t, "new-slug-2", newCmt.Slug)
		assert.Equal(t, "new-author-2", newCmt.Author)
		assert.Equal(t, "new-body-2", newCmt.Body)
	})

	t.Run("test get comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "new-author",
			Body:   "new-body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		assert.Equal(t, "new-slug", newCmt.Slug)
		assert.Equal(t, "new-author", newCmt.Author)
		assert.Equal(t, "new-body", newCmt.Body)
	})
}
