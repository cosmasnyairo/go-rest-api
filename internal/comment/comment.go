package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("failed to fetch comment by id")
	ErrorNotImplemented  = errors.New("not implemented")
)

// Defines all the methods our service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Comment Body
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Struct where all our logic will be built on top of
type Service struct {
	Store Store
}

// Returns pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrorNotImplemented
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorFetchingComment
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrorNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrorNotImplemented
}
