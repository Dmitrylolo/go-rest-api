package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment a representation of a comment structure
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store is an interface that defines the
// methods that service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service is a struct on which all
// the business logic of the service
type Service struct {
	Store Store
}

// NewService rerturns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment returns a comment by its id from the store
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving comment with id: ", id)

	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println("Error retrieving comment: ", err)
		return Comment{}, ErrFetchingComment
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, updatedComment Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, id, updatedComment)
	if err != nil {
		fmt.Println("Error updating comment")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	createdComment, err := s.Store.CreateComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return createdComment, nil
}
