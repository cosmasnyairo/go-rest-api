package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorCreatingComment = errors.New("failed to create comment")
	ErrorDeletingComment = "failed to delete comment"
	ErrorUpdatingComment = "failed to update comment"
	ErrorFetchingComment = errors.New("failed to fetch comment by id")
	ErrorNotImplemented  = errors.New("not implemented")
	EmptyComment         = Comment{}
)

// Defines all the methods our service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error)
}

// Comment Body required by our service
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
	fmt.Println("Creating comment")
	comment, err := s.Store.CreateComment(ctx, cmt)
	if err != nil {
		return EmptyComment, fmt.Errorf("%s: %w", ErrorCreatingComment, err)
	}
	return comment, ErrorNotImplemented
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving comment")
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		return EmptyComment, fmt.Errorf("%s: %w", ErrorFetchingComment, err)
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	fmt.Println("Updating comment")
	updatedcomment, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		return EmptyComment, fmt.Errorf("%s: %w", ErrorUpdatingComment, err)
	}
	return updatedcomment, err
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	fmt.Println("Deleting comment")
	return s.Store.DeleteComment(ctx, id)
}
