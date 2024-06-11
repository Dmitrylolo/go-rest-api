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

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
